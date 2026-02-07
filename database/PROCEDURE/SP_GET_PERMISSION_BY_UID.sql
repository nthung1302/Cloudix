CREATE OR ALTER PROCEDURE SP_GET_PERMISSION_BY_UID
    @UID UNIQUEIDENTIFIER
AS
BEGIN
    SET NOCOUNT ON;
    SET XACT_ABORT ON;
	-- Results: Status, Message, PermissionGroupName, PermissionName
    BEGIN TRY

        DECLARE @AccountID INT;

        SELECT @AccountID = AccountID FROM Accounts WHERE UID = @UID;
        IF @AccountID IS NULL
        BEGIN
            SELECT 0 AS Status, N'Tài khoản không tồn tại.' AS Message;
            RETURN;
        END

        SELECT
            pg.PermissionGroupName,
            p.PermissionName
        FROM Accounts a
        INNER JOIN AccountPermissions ap ON a.AccountID = ap.AccountID
        INNER JOIN Permissions p ON ap.PermissionID = p.PermissionID
        LEFT JOIN PermissionGroups pg ON p.PermissionGroupID = pg.PermissionGroupID
        WHERE a.UID = @UID;

        SELECT 1 AS Status, N'Lấy danh sách quyền thành công.' AS Message;
    END TRY
    
    BEGIN CATCH
        IF @@TRANCOUNT > 0 ROLLBACK TRANSACTION;
        SELECT  -1 AS Status, ERROR_MESSAGE() AS Message;
    END CATCH
END
-- EXEC SP_GET_ACCOUNT_PERMISSION @UID = '24D5C6CF-1090-4DAF-974B-2603C35426A0'