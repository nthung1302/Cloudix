-- =====================================
-- Author: Cloudix Dev Team
-- Create date: 2024-06-15
-- Description: Table 
-- =====================================

CREATE DATABASE Cloudix;
GO
USE Cloudix;
GO

-- =====================================
-- TABLE Permissions
-- =====================================
CREATE TABLE PermissionGroup (
    PermissionGroupID INT IDENTITY(1,1) PRIMARY KEY,
    PermissionGroupName NVARCHAR(100) NOT NULL UNIQUE,
    Description NVARCHAR(255)
);

CREATE TABLE Permissions (
    PermissionID INT IDENTITY(1,1) PRIMARY KEY,
    PermissionName NVARCHAR(100) NOT NULL UNIQUE,
    Description NVARCHAR(255),
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE(),
    PermissionGroupID INT,
    CONSTRAINT FK_Permissions_PermissionGroupID FOREIGN KEY (PermissionGroupID) REFERENCES PermissionGroup(PermissionGroupID) ON DELETE SET NULL
);
CREATE INDEX IDX_Permissions_PermissionName ON Permissions(PermissionName);

-- =====================================
-- TABLE Roles
-- =====================================
CREATE TABLE Roles (
    RoleID INT IDENTITY(1,1) PRIMARY KEY,
    RoleName NVARCHAR(100) NOT NULL UNIQUE,
    Description NVARCHAR(255),
    IsDefault BIT DEFAULT 0
);
CREATE INDEX IDX_Roles_RoleName ON Roles(RoleName);

-- =====================================
-- TABLE Accounts
-- =====================================
CREATE TABLE Accounts (
    AccountID INT IDENTITY(1,1) PRIMARY KEY,
    UID UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),  
    UserName NVARCHAR(50) NOT NULL UNIQUE,
    Password NVARCHAR(255) NOT NULL,
    IsActive BIT DEFAULT 1,
    ExpiredAt DATETIME
);
CREATE INDEX IDX_Accounts_UserName ON Accounts(UserName);

-- =====================================
-- TABLE Role/Account Permissions
-- =====================================
CREATE TABLE RolePermissions (
    RoleID INT NOT NULL,
    PermissionID INT NOT NULL,
    PRIMARY KEY (RoleID, PermissionID),
    CONSTRAINT FK_RolePermissions_RoleID FOREIGN KEY (RoleID) REFERENCES Roles(RoleID) ON DELETE CASCADE,
    CONSTRAINT FK_RolePermissions_PermissionID FOREIGN KEY (PermissionID) REFERENCES Permissions(PermissionID) ON DELETE CASCADE
);

CREATE INDEX IDX_RolePermissions_RoleID ON RolePermissions(RoleID);
CREATE INDEX IDX_RolePermissions_PermissionID ON RolePermissions(PermissionID);

CREATE TABLE AccountPermissions (
    AccountID INT NOT NULL,
    PermissionID INT NOT NULL,
    CONSTRAINT PK_AccountPermissions PRIMARY KEY (AccountID, PermissionID),
    CONSTRAINT FK_AccountPermissions_AccountID FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE CASCADE,
    CONSTRAINT FK_AccountPermissions_PermissionID FOREIGN KEY (PermissionID) REFERENCES Permissions(PermissionID) ON DELETE CASCADE
);
CREATE INDEX IDX_AccountPermissions_AccountID ON AccountPermissions(AccountID);
CREATE INDEX IDX_AccountPermissions_PermissionID ON AccountPermissions(PermissionID);
-- =====================================
-- TABLE AccountRoles
-- =====================================
CREATE TABLE AccountRoles (
    AccountID INT NOT NULL,
    RoleID INT NOT NULL,
    PRIMARY KEY (AccountID, RoleID),
    CONSTRAINT FK_AccountRoles_AccountID FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE CASCADE,
    CONSTRAINT FK_AccountRoles_RoleID FOREIGN KEY (RoleID) REFERENCES Roles(RoleID) ON DELETE CASCADE
);

-- =====================================
-- TABLE Profiles
-- =====================================
CREATE TABLE AccountProfiles (
    AccountID INT PRIMARY KEY,
    FullName NVARCHAR(100),
    Email NVARCHAR(100) UNIQUE,
    PhoneNumber NVARCHAR(15),
    Address NVARCHAR(255),
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE(),
    CONSTRAINT FK_Profiles_AccountID FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE CASCADE
);

-- =====================================
-- TABLE AuditLogs
-- =====================================
CREATE TABLE AuditLogs (
    LogID INT IDENTITY(1,1) PRIMARY KEY,
    AccountID INT,
    Action NVARCHAR(100) NOT NULL,
    Timestamp DATETIME DEFAULT GETDATE(),
    Details NVARCHAR(1000),
    CONSTRAINT FK_AuditLogs_AccountID FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE SET NULL
);

-- =====================================
-- TABLE ACCOUNT SETTINGS
-- =====================================


INSERT INTO Roles (RoleName, Description) VALUES
('admin', N'Quyền quản trị viên'),
('administrator', N'Quyền quản lý hệ thống'),
('user', N'Quyền người dùng bình thường');

INSERT INTO PermissionGroup (PermissionGroupName, Description) VALUES
('MANAGER_USER', N'Quản lý người dùng')

INSERT INTO Permissions (PermissionName, Description, PermissionGroupID) VALUES
('CREATE_USER', N'Tạo người dùng mới', 1),
('DELETE_USER', N'Xóa người dùng', 1),
('UPDATE_USER', N'Cập nhật thông tin người dùng', 1),
('VIEW_USER', N'Xem thông tin người dùng', 1);

SELECT p.PermissionName, p.Description
FROM Permissions p 
LEFT JOIN PermissionGroup pg ON pg.PermissionGroupID = p.PermissionGroupID
INNER JOIN AccountPermissions ap ON p.PermissionID = ap.PermissionID
INNER JOIN Accounts a ON ap.AccountID = a.AccountID
WHERE a.UserName = 'admin';