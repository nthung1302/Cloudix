/* ================================
   DATABASE
================================ */
CREATE DATABASE Cloudix_Server
GO
USE Cloudix_Server
GO

/* ================================
   ROLE GROUP
================================ */
CREATE TABLE RoleGroup (
    id INT IDENTITY(1,1) PRIMARY KEY,
    code VARCHAR(30) NOT NULL UNIQUE,
    description NVARCHAR(100) NOT NULL
)

/* ================================
   PERMISSION
================================ */
CREATE TABLE Permission (
    id INT IDENTITY(1,1) PRIMARY KEY,
    code VARCHAR(30) NOT NULL UNIQUE,
    description NVARCHAR(100) NOT NULL,
    is_default BIT DEFAULT 0,

    role_group_id INT,
    CONSTRAINT fk_permission_role_group
        FOREIGN KEY (role_group_id) REFERENCES RoleGroup(id)
)

CREATE INDEX idx_permission ON Permission (is_default, code, description);

/* ================================
   ACCOUNT
================================ */
CREATE TABLE Account (
    id INT IDENTITY(1,1) PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(10) DEFAULT 'user',
    is_active BIT DEFAULT 0,
    is_locked BIT DEFAULT 0,
    expired_at DATETIME2 DEFAULT GETDATE()
)

CREATE INDEX idx_account ON Account (is_active, is_locked, expired_at);

/* ================================
   ACCOUNT - PERMISSION
================================ */
CREATE TABLE AccountPermission (
    permission_id INT,
    account_id INT,

    CONSTRAINT pk_account_permission PRIMARY KEY (permission_id, account_id),
    CONSTRAINT fk_ap_permission FOREIGN KEY (permission_id) REFERENCES Permission(id),
    CONSTRAINT fk_ap_account FOREIGN KEY (account_id) REFERENCES Account(id)
)

/* ================================
   USER PROFILE
================================ */
CREATE TABLE UserProfile (
    id INT IDENTITY(1,1) PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(12),
    full_name NVARCHAR(50),
    address NVARCHAR(100),

    created_at DATETIME2 DEFAULT GETDATE(),
    updated_at DATETIME2 DEFAULT GETDATE(),

    account_id INT, 
	CONSTRAINT fk_user_profile_account FOREIGN KEY (account_id) REFERENCES Account(id)
)


USE [Cloudix_Server]
GO
/****** Object:  StoredProcedure [dbo].[SP_CREATE_ACCOUNT]    Script Date: 04/02/2026 12:00:18 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
ALTER PROCEDURE [dbo].[SP_CREATE_ACCOUNT]
    @username   VARCHAR(20),
    @password   VARCHAR(255),
    @full_name  NVARCHAR(50),
    @email      VARCHAR(50)
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;
	BEGIN TRY
		IF EXISTS (SELECT 1 FROM Account WHERE username = @username)
		BEGIN
			SELECT 0 AS status, N'Username đã tồn tại' AS message;
			RETURN;
		END

		-- Email tồn tại
		IF EXISTS (SELECT 1 FROM UserProfile WHERE email = @email)
		BEGIN
			SELECT 0 AS status, N'Email đã tồn tại' AS message;
			RETURN;
		END

		BEGIN TRANSACTION;

		INSERT INTO Account (username, [password], is_active, is_locked)
		VALUES (@username, @password, 1, 0);

		DECLARE @account_id INT = SCOPE_IDENTITY();

		INSERT INTO UserProfile (email, full_name, account_id)
		VALUES (@email, @full_name, @account_id);

		INSERT INTO AccountPermission (permission_id, account_id)
		SELECT id, @account_id FROM Permission WHERE is_default = 1;

		COMMIT TRANSACTION;

		SELECT 1 AS status, N'Đăng ký thành công' AS message;
	END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;
        THROW;
    END CATCH
END
-- EXEC SP_CREATE_ACCOUNT @username='admin2', @password='sdsd', @full_name='Nguyễn Văn B', @email='admin2@gmail.com'

ALTER PROCEDURE [dbo].[SP_LOGIN_ACCOUNT]
    @username VARCHAR(50)
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @account_id INT;

    -- Email hoặc username
    IF CHARINDEX('@', @username) > 0
        SELECT @account_id = u.account_id FROM UserProfile u WHERE u.email = @username;
    ELSE
        SELECT @account_id = a.id FROM Account a WHERE a.username = @username;

    IF @account_id IS NULL
    BEGIN
        SELECT 0 AS status, N'Tên đăng nhập không tồn tại' AS message;
        RETURN;
    END

    IF EXISTS (SELECT 1 FROM Account WHERE id = @account_id AND is_locked = 1)
    BEGIN
        SELECT 0 AS status, N'Tài khoản bị khóa' AS message;
        RETURN;
    END

    IF EXISTS (SELECT 1 FROM Account WHERE id = @account_id AND is_active = 0)
    BEGIN
        SELECT 0 AS status, N'Tài khoản chưa kích hoạt' AS message;
        RETURN;
    END

    SELECT
        1 AS status,
        N'OK' AS message,
        id,
        username,
        password,
        is_active,
        is_locked,
        expired_at
    FROM Account
    WHERE id = @account_id;
END
