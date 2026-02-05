CREATE OR ALTER PROCEDURE SP_CHANGE_PASSWORD
	@UID UNIQUEIDENTIFIER,
	@NewPassword VARCHAR(100)
AS
BEGIN
	SET NOCOUNT ON;
	SET XACT_ABORT ON;
	DECLARE @error INT;
	DECLARE @message NVARCHAR(255);

	-- Kiểm tra tài khoản tồn tại không
	IF NOT EXISTS (SELECT 1 FROM Account WHERE UID = @UID)
    BEGIN
		set @error = -1;
		set @message = N'Account does not exist.';
        SELECT @error AS Error, @message AS Message;
        RETURN;
    END

	-- Cập nhật mật khẩu mới
	UPDATE Account set password = @NewPassword WHERE UID = @UID

	SET @error = 0;
	SET @message = N'Password has been changed successfully.';
	SELECT @error AS Error, @message AS Message;
END