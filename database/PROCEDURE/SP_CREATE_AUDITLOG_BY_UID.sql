CREATE OR ALTER PROCEDURE SP_CREATE_AUDITLOG_BY_UID
    @UID UNIQUEIDENTIFIER,
    @Action NVARCHAR(100),
    @Details NVARCHAR(500)
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        DECLARE @AccountID INT;

        SELECT @AccountID = AccountID FROM Accounts WHERE UID = @UID;

        IF @AccountID IS NULL
        BEGIN
            ROLLBACK TRANSACTION;
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message;
            RETURN;
        END

        INSERT INTO AuditLogs (AccountID, Action, TimeStamp, Details)
        VALUES (@AccountID, @Action, GETDATE(), @Details);

        COMMIT TRANSACTION;
        SELECT 1 AS Status, N'Nhật ký đã được tạo thành công.' AS Message;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT -1 AS Status, ERROR_MESSAGE() AS Message;
    END CATCH
END
-- EXEC SP_CREATE_AUDITLOG_BY_UID @UID = '24D5C6CF-1090-4DAF-974B-2603C35426A0', @Action = 'Login', @Details = 'User logged in successfully.'