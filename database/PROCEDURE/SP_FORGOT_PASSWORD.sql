CREATE OR ALTER PROCEDURE SP_FORGOT_PASSWORD
    @UserName NVARCHAR(50),
    @NewPassword NVARCHAR(255)
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;
    
    -- Results: Status, Message
    BEGIN TRY
        BEGIN TRANSACTION;

        DECLARE @AccountID INT;

        IF CHARINDEX('@', @UserName) > 0
            SELECT @AccountID = AccountID FROM Profiles WHERE Email = @UserName;
        ELSE
            SELECT @AccountID = AccountID FROM Accounts WHERE UserName = @UserName;

        IF @AccountID IS NULL
        BEGIN
            ROLLBACK TRANSACTION;
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message;
            RETURN;
        END

        IF EXISTS (SELECT 1 FROM Accounts WHERE AccountID = @AccountID AND IsActive = 0)
        BEGIN
            ROLLBACK TRANSACTION;
            SELECT 0 AS Status, N'Tài khoản đã bị khóa.' AS Message;
            RETURN;
        END

        UPDATE Accounts
        SET Password = @NewPassword
        WHERE AccountID = @AccountID;

        COMMIT TRANSACTION;
        SELECT 1 AS Status, N'Mật khẩu đã được đặt lại thành công.' AS Message;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT -1 AS Status, ERROR_MESSAGE() AS Message;
    END CATCH
END
-- EXEC SP_FORGOT_PASSWORD @UserName = 'admin', @NewPassword = 'newpassword123'