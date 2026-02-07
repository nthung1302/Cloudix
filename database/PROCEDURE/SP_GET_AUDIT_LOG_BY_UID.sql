CREATE OR ALTER PROCEDURE SP_GET_AUDIT_LOG_BY_UID
    @UID UNIQUEIDENTIFIER
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;
    BEGIN TRY
		DECLARE @AccountID INT;

        SELECT @AccountID = AccountID FROM Accounts WHERE UID = @UID AND IsActive = 1;
        
        IF @AccountID IS NULL
        BEGIN
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message;
            RETURN;
        END

        SELECT LogID, AccountID, Action, TimeStamp, Details
        FROM AuditLogs
        WHERE AccountID = @AccountID ORDER BY TimeStamp DESC;

        SELECT 1 AS Status, N'Lấy nhật ký thành công.' AS Message
        
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT -1 AS Status, ERROR_MESSAGE() AS Message;
    END CATCH
END
-- EXEC SP_GET_AUDIT_LOG_BY_UID @UID = '24D5C6CF-1090-4DAF-974B-2603C35426A0'