CREATE OR ALTER PROCEDURE SP_UPDATE_PASWORD_BY_UID
    @UID UNIQUEIDENTIFIER,
    @NewPassword VARCHAR(255)
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        -- Kiểm tra tài khoản tồn tại
        IF NOT EXISTS (SELECT 1 FROM Accounts WHERE UID = @UID)
        BEGIN
            ROLLBACK TRANSACTION;
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message;
            RETURN;
        END

        -- Kiểm tra tài khoản bị khóa
        IF EXISTS (SELECT 1 FROM Accounts WHERE UID = @UID AND IsActive = 0)
        BEGIN
            ROLLBACK TRANSACTION;
            SELECT 0 AS Status, N'Tài khoản đã bị khóa.' AS Message;
            RETURN;
        END

        -- Cập nhật mật khẩu
        UPDATE Accounts
        SET Password = @NewPassword
        WHERE UID = @UID;

        COMMIT TRANSACTION;
        SELECT 1 AS Status, N'Mật khẩu đã được thay đổi thành công.' AS Message;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT -1 AS Status, ERROR_MESSAGE() AS Message;
    END CATCH
END
-- EXEC SP_CHANGE_PASSWORD @UID = 'some-uid-here', @NewPassword = 'newpassword123'