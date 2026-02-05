CREATE OR ALTER PROCEDURE SP_GET_ACCOUNT_INFO
    @UID UNIQUEIDENTIFIER
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;

    -- Results: Status, Message, UserName, FullName, Email, CreatedAt
    BEGIN TRY

        DECLARE @AccountID INT;

        DECLARE @UserName VARCHAR(50);
        DECLARE @FullName NVARCHAR(100);
        DECLARE @Email NVARCHAR(100);
		DECLARE @CreatedAt DATETIME;

        SELECT @AccountID = AccountID FROM Accounts WHERE UID = @UID;
        IF @AccountID IS NULL
        BEGIN
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message, NULL AS UserName, NULL AS FullName, NULL AS Email;
            RETURN;
        END

        SET @UserName = (SELECT UserName FROM Accounts WHERE UID = @UID);
        SET @FullName = (SELECT FullName FROM AccountProfiles WHERE AccountID = @AccountID);
        SET @Email = (SELECT Email FROM AccountProfiles WHERE AccountID = @AccountID);
        SET @CreatedAt = (SELECT CAST(CreatedAt AS DATE) FROM AccountProfiles WHERE AccountID = @AccountID);

        SELECT 1 AS Status, N'Lấy thông tin tài khoản thành công.' AS Message, @UserName AS UserName, @FullName AS FullName, @Email AS Email, @CreatedAt AS CreatedAt;
    END TRY
    
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT  -1 AS Status, ERROR_MESSAGE() AS Message;
    END CATCH
END
-- EXEC SP_GET_ACCOUNT_INFO @UID = '24D5C6CF-1090-4DAF-974B-2603C35426A0'