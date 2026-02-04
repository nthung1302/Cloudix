CREATE OR ALTER PROCEDURE [dbo].[SP_LOGIN_ACCOUNT]
    @username VARCHAR(50)
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @account_id INT;

    IF CHARINDEX('@', @username) > 0
        SELECT @account_id = u.account_id FROM UserProfile u WHERE u.email = @username;
    ELSE
        SELECT @account_id = a.id FROM Account a WHERE a.username = @username;

    IF @account_id IS NULL
    BEGIN
        SELECT 0 AS status, N'Tên đăng nhập không tồn tại' AS message, NULL AS id, NULL AS username, NULL AS password, NULL AS expired_at;
        RETURN;
    END

    IF EXISTS (SELECT 1 FROM Account WHERE id = @account_id AND is_locked = 1)
    BEGIN
        SELECT 0 AS status, N'Tài khoản bị khóa' AS message, NULL AS id, NULL AS username, NULL AS password, NULL AS expired_at;
        RETURN;
    END

    IF EXISTS (SELECT 1 FROM Account WHERE id = @account_id AND is_active = 0)
    BEGIN
        SELECT 0 AS status, N'Tài khoản chưa kích hoạt' AS message, NULL AS id, NULL AS username, NULL AS password, NULL AS expired_at;
        RETURN;
    END
	
    SELECT 1 AS status, N'OK' AS message, id, username, password, expired_at
    FROM Account WHERE id = @account_id;
END
-- EXEC SP_LOGIN_ACCOUNT @username = 'admi'