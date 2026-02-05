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


select * from Account

SELECT CAST(expired_at AS DATE) AS ExpiredDate from Account