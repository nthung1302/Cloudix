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
CREATE TABLE Permissions (
    PermissionID INT IDENTITY(1,1) PRIMARY KEY,
    PermissionName NVARCHAR(100) NOT NULL UNIQUE,
    Description NVARCHAR(255)
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
    FOREIGN KEY (RoleID) REFERENCES Roles(RoleID) ON DELETE CASCADE,
    FOREIGN KEY (PermissionID) REFERENCES Permissions(PermissionID) ON DELETE CASCADE
);

CREATE INDEX IDX_RolePermissions_RoleID ON RolePermissions(RoleID);
CREATE INDEX IDX_RolePermissions_PermissionID ON RolePermissions(PermissionID);

CREATE TABLE AccountPermissions (
    AccountID INT NOT NULL,
    PermissionID INT NOT NULL,
    PRIMARY KEY (AccountID, PermissionID),
    FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE CASCADE,
    FOREIGN KEY (PermissionID) REFERENCES Permissions(PermissionID) ON DELETE CASCADE
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
    FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE CASCADE,
    FOREIGN KEY (RoleID) REFERENCES Roles(RoleID) ON DELETE CASCADE
);

-- =====================================
-- TABLE Profiles
-- =====================================
CREATE TABLE Profiles (
    AccountID INT PRIMARY KEY,
    FullName NVARCHAR(100),
    Email NVARCHAR(100) UNIQUE,
    PhoneNumber NVARCHAR(15),
    Address NVARCHAR(255),
    FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE CASCADE,
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE()
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
    FOREIGN KEY (AccountID) REFERENCES Accounts(AccountID) ON DELETE SET NULL
);

-- =====================================
-- TABLE ACCOUNT SETTINGS
-- =====================================


-- =====================================
-- STORED PROCEDURES
-- =====================================
