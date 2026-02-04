CREATE OR ALTER PROCEDURE SP_CREATE_ACCOUNT
    @username   VARCHAR(20),
    @password   VARCHAR(255),
    @full_name  NVARCHAR(50),
    @email      VARCHAR(50)
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;
	BEGIN TRY
		IF EXISTS (SELECT 1 FROM Account WHERE username = @username)
		BEGIN
			SELECT 0 AS status, N'Username đã tồn tại' AS message;
			RETURN;
		END

		-- Email tồn tại
		IF EXISTS (SELECT 1 FROM UserProfile WHERE email = @email)
		BEGIN
			SELECT 0 AS status, N'Email đã tồn tại' AS message;
			RETURN;
		END

		BEGIN TRANSACTION;

		INSERT INTO Account (username, [password], is_active, is_locked)
		VALUES (@username, @password, 1, 0);

		DECLARE @account_id INT = SCOPE_IDENTITY();

		INSERT INTO UserProfile (email, full_name, account_id)
		VALUES (@email, @full_name, @account_id);

		INSERT INTO AccountPermission (permission_id, account_id)
		SELECT id, @account_id FROM Permission WHERE is_default = 1;

		COMMIT TRANSACTION;

		SELECT 1 AS status, N'Đăng ký thành công' AS message;
	END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;
        THROW;
    END CATCH
END
-- EXEC SP_CREATE_ACCOUNT @username='admin2', @password='sdsd', @full_name='Nguyễn Văn B', @email='admin2@gmail.com'