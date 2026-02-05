CREATE OR ALTER PROCEDURE [dbo].[SP_LOGIN_ACCOUNT]
    @UserName VARCHAR(50)
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @AccountID INT
	DECLARE @Error INT
	DECLARE @Message NVARCHAR(255)
	DECLARE @UID VARCHAR(100)
	DECLARE @Password VARCHAR(100)
	DECLARE @ExpriedAt DATE

    IF CHARINDEX('@', @UserName) > 0
        SELECT @AccountID = u.AccountID FROM Profiles u WHERE u.Email = @UserName;
    ELSE
        SELECT @AccountID = a.AccountID FROM Accounts a WHERE a.UserName = @UserName;

    IF @AccountID IS NULL
    BEGIN
        SET @Error = -1;
        SET @Message = N'Account does not exist.';
        SELECT @Error AS Error, @Message AS Message, NULL AS id, NULL AS username, NULL AS password, NULL AS expired_at;
        RETURN;
    END

    IF EXISTS (SELECT 1 FROM Accounts WHERE AccountID = @AccountID AND IsActive = 0)
    BEGIN
        SET @Error = -1;
        SET @Message = N'Account is inactive.';
        SELECT @Error AS Error, @Message AS Message, NULL AS id, NULL AS username, NULL AS password, NULL AS expired_at;
        RETURN;
    END
	
    SELECT @UID = UID, @Password = Password, @ExpriedAt = ExpiredAt
    FROM Accounts
    WHERE AccountID = @AccountID;

    SET @Error = 0;
    SET @Message = N'Login successful.';
    SELECT @Error AS Error, @Message AS Message, @UID AS id, @UserName AS username, @Password AS password, @ExpriedAt AS expired_at;
END
-- EXEC SP_LOGIN_ACCOUNT @username = 'admi'