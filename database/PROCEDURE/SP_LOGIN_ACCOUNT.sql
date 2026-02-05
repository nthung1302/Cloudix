CREATE OR ALTER PROCEDURE SP_LOGIN_ACCOUNT
    @UserName VARCHAR(50)
AS
BEGIN
    -- Results: Status, Message, UID, UserName, Password, ExpiredAt
    SET NOCOUNT ON;

    BEGIN TRY
        DECLARE @AccountID INT
        DECLARE @UID VARCHAR(100)
        DECLARE @Password VARCHAR(100)
        DECLARE @ExpiredAt DATETIME

        IF CHARINDEX('@', @UserName) > 0
            SELECT @AccountID = AccountID FROM AccountProfiles WHERE Email = @UserName;
        ELSE
            SELECT @AccountID = AccountID FROM Accounts WHERE UserName = @UserName;

        IF @AccountID IS NULL
        BEGIN
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message,
                   NULL AS UID, NULL AS UserName, NULL AS Password, NULL AS ExpiredAt;
            RETURN;
        END

        IF EXISTS (SELECT 1 FROM Accounts WHERE AccountID = @AccountID AND IsActive = 0)
        BEGIN
            SELECT 0 AS Status, N'Tài khoản đã bị khóa.' AS Message,
                   NULL AS UID, NULL AS UserName, NULL AS Password, NULL AS ExpiredAt;
            RETURN;
        END

        SELECT  @UID = UID, @Password = Password, @ExpiredAt = ExpiredAt
        FROM Accounts
        WHERE AccountID = @AccountID;

        SELECT 1 AS Status, N'Đăng nhập thành công.' AS Message, @UID AS UID, @UserName AS UserName, @Password AS Password, @ExpiredAt AS ExpiredAt;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT -1 AS Status, ERROR_MESSAGE() AS Message, NULL AS UID, NULL AS UserName, NULL AS Password, NULL AS ExpiredAt;
    END CATCH
END
-- EXEC SP_LOGIN_ACCOUNT @username = 'admi'
