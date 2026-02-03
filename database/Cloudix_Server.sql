CREATE DATABASE Cloudix_Server
USE Cloudix_Server
GO

CREATE TABLE NhomQuyen (
	id int identity(1,1) primary key,
	gia_tri varchar(30) not null unique,
	mo_ta NVARCHAR(100) not null
)

CREATE TABLE TriTietQuyen (
	id int identity(1,1) primary key,
	gia_tri varchar(30) not null unique,
	mo_ta NVARCHAR(100) not null,
	mac_dinh BIT DEFAULT 0,

	nhom_quyen_id int, CONSTRAINT fk_ttq_nhom_quyen FOREIGN KEY (nhom_quyen_id) REFERENCES NhomQuyen (id)
)
CREATE INDEX idx_tri_tiet_quyen ON TriTietQuyen (mac_dinh, gia_tri, mo_ta);

CREATE TABLE TaiKhoan (
	id int identity(1,1) primary key,
	ten_dang_nhap varchar(20) not null unique,
	mat_khau varchar(100) not null,
	quyen varchar(5) default 'user',
	hoat_dong bit default 0,
	tk_cam bit default 0,
	ngay_het_han DATETIME2 DEFAULT GETDATE()
)

CREATE INDEX idx_tai_khoan ON TaiKhoan(hoat_dong, tk_cam, ngay_het_han);

CREATE TABLE QuyenTaiKhoan (
	ma_tritietquyen int,
	ma_taikhoan int,

	constraint pk_qtk primary key (ma_tritietquyen, ma_taikhoan),
	constraint fk_qtk_tritietquyen foreign key (ma_tritietquyen) references TriTietQuyen(id),
	constraint fk_qtk_taikhoan foreign key (ma_taikhoan) references TaiKhoan(id)
)

CREATE TABLE HoSoNguoiDung (
	id int identity(1,1) primary key,
	email varchar(20) unique,
	dien_thoai varchar(12) unique,
	ten_day_du NVARCHAR(50),
	dia_chi NVARCHAR(100),
	
	ngay_tao DATETIME2 DEFAULT GETDATE(),
	ngay_sua DATETIME2 DEFAULT GETDATE(),
	
	ma_taikhoan int,
	constraint fk_hsnd_taikhoan foreign key (ma_taikhoan) references TaiKhoan(id)
)

