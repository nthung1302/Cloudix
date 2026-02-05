CREATE OR ALTER PROCEDURE SP_CREATE_ACCOUNT
    @UserName NVARCHAR(50),
    @Password NVARCHAR(255),
    @FullName NVARCHAR(100),
    @Email NVARCHAR(100)
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;

    DECLARE @error INT;
    DECLARE @message NVARCHAR(255);
    
    BEGIN TRY
        BEGIN TRANSACTION;

        -- Kiểm tra trùng username và email
        IF EXISTS (SELECT 1 FROM Accounts WHERE UserName = @UserName)
        BEGIN
            SET @error = -1;
            SET @message = 'Username already exists.';
            ROLLBACK TRANSACTION;
            SELECT @error AS Error, @message AS Message;
            RETURN;
        END

        IF EXISTS (SELECT 1 FROM Profiles WHERE Email = @Email)
        BEGIN
            SET @error = -2;
            SET @message = 'Email already exists.';
            ROLLBACK TRANSACTION;
            SELECT @error AS Error, @message AS Message;
            RETURN;
        END

        -- Tạo tài khoản mới
        INSERT INTO Accounts (UserName, Password)
        VALUES (@UserName, @Password);

        -- Lấy AccountID vừa tạo
        DECLARE @AccountID INT = SCOPE_IDENTITY();

		-- Cập nhật ngày hết hạn 
		UPDATE Accounts SET ExpiredAt = DATEADD(YEAR, 1, GETDATE()) 
		WHERE AccountID = @AccountID;

        -- Tạo profile cho tài khoản
        INSERT INTO Profiles (AccountID, FullName, Email)
        VALUES (@AccountID, @FullName, @Email);

        -- Gán vai trò và quyền mặc định
        INSERT INTO AccountRoles (AccountID, RoleID)
        SELECT @AccountID, RoleID
        FROM Roles
        WHERE IsDefault = 1;

        -- Gán quyền mặc định từ vai trò
        INSERT INTO AccountPermissions (AccountID, PermissionID)
        SELECT @AccountID, rp.PermissionID
        FROM RolePermissions rp
        JOIN Roles r ON rp.RoleID = r.RoleID
        WHERE r.IsDefault = 1;

        COMMIT TRANSACTION;
        SET @error = 0;
        SET @message = 'Account created successfully.';
        SELECT @error AS Error, @message AS Message;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;
        SET @error = -99;
        SET @message = ERROR_MESSAGE();
        SELECT @error AS Error, @message AS Message;
    END CATCH
END
-- EXEC SP_CREATE_ACCOUNT @UserName='admin', @Password='admin@123', @FullName='Nguyễn Minh Hùng', @Email='admin@example.com'
-- SELECT * FROM Accounts