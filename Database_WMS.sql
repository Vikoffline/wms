drop database if exists WMS;
create database if not exists WMS;
use WMS;



create table Instances (
	IdNum int Unique Key auto_increment,
    Id varchar(16) primary key default("In_"),
    Type enum("Склад", "СЦ", "ПВЗ"),
	Coordinates varchar(32),
    Capacity int,
    IsAvailable bool
);

create table instancesInfo (
	IdNum int Unique Key auto_increment,
	instanceId varchar(16) primary key unique,
    foreign key (instanceId) references Instances(Id) on delete cascade, 
    ContactNumber varchar(12) default "",
    Email varchar(100) default "",
    WorkingHours varchar(50) default "",
    Length int default 0,
    Width int default 0,
    Height int default 0,
    Volume int default(Length*Width*Height),
    City varchar(25) default "",
	Adress varchar(96) default ""
);

create table instanceParts (
	IdNum int Unique Key auto_increment,
    Id varchar(16) primary key default("iP_"),
    Type enum("Офис", "Склад", "Спец. зона"),
    itemMaxSize int,
    Capacity int,
    instanceId varchar(16),
    foreign key (instanceId) references Instances(Id) on delete cascade
);

create table Items (
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("It_"),
    Size int,
    vendorId varchar(16),
    Name varchar(16)
);

create table Permissions(
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("Pr_"),
    Code int unique,
    Name varchar(16),
    tableName varchar(16)
);

create table Roles (
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("Rl_"),
    Name varchar(16)
);

create table Roles_Perms (
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("RP_"),
    roleId varchar(16),
    permId varchar(16),
    foreign key (roleId) references Roles(Id) on delete cascade,
    foreign key (permId) references Permissions(Id) on delete cascade,
    unique(roleId, permId)
);

create table Managers(
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("Mn_"),
    Login VARCHAR(16) unique,
    Password VARCHAR(64),
    Name varchar(32),
    ContactNumber varchar(12),
    Email varchar(100),
    roleId varchar(16),
    foreign key (roleId) references Roles(Id) on delete set null
);

create table Actions (
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("Ac_"),
	Type enum("Связь", "Траспортировка", "Хранение"),
    Date datetime default(current_date()), 
    itemId varchar(16),
    instId varchar(16),
    managerId varchar(16), -- УБрать атрибут
    foreign key (itemId) references Items(Id) on delete cascade,
    foreign key (instId) references Instances(Id) on delete cascade,
    foreign key (managerId) references Managers(Id) on delete set null
);

create table Sessions (
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("Sn_"),
    Token char(64) unique,
    managerId varchar(16),
    foreign key (managerId) references Managers(Id) on delete set null
);

create table Logging (
	IdNum int auto_increment Unique Key,
	Id varchar(16) primary key default("Lg_"),
    funcName varchar(32), 
    Date datetime,
    errText text
);


delimiter //

create trigger autoInfo after insert on Instances for each row begin
	insert into instancesInfo(instanceId) values (new.Id);
end//

create trigger instancesBeutyId before insert on Instances for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Instances T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger instancePartsBeutyId before insert on instanceParts for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from instanceParts T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger itemsBeutyId before insert on Items for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Items T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger roles_permsBeutyId before insert on Roles_Perms for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Roles_Perms T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger rolesBeutyId before insert on Roles for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Roles T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger permissionsBeutyId before insert on Permissions for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Permissions T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger managersBeutyId before insert on Managers for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Managers T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger actionsBeutyId before insert on Actions for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Actions T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger sessionsBeutyId before insert on Sessions for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Sessions T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

create trigger loggingBeutyId before insert on Logging for each row begin
	declare IdNum int default 0;
    select count(*) into IdNum from Logging T;
	if (new.Id like "__!_" escape "!") then
		set new.Id = concat(new.Id, IdNum+1);
	end if;
end//

delimiter ;

drop user if exists site;
create user site identified by "DevPassword";
grant all on wms.* to site; 

insert into Permissions(Code, Name, tableName) values
(0, "AllRights", "AllTables"),
(1, "get", "Managers");


insert into Roles(Name) values ("root");
insert into Roles(Name) values ("empty");
insert into Roles(Name) values ("test");

insert into Roles_Perms(roleId, permId) values ("Rl_1", "Pr_1");
insert into Roles_Perms(roleId, permId) values ("Rl_3", "Pr_2");

insert into Managers(Login, Password, Name, ContactNumber, Email, roleId) values ("root", "$2a$10$7z2Qu0bttRd2T3ea0Fzluu1Lp8iyU2sStJByuhYBQhE3hKENWe2Tm", "", "", "", "Rl_1");

select * from Managers;
select * from Logging where IdNum >= 1 order by IdNum desc limit 20;


select Sn.Id from Sessions Sn inner join Managers Mn on Mn.Id = Sn.managerId inner join Roles Rl on Rl.Id = Mn.roleId inner join Roles_Perms RP on RP.roleId = Rl.Id inner join Permissions Pr on Pr.Id = RP.permId where Sn.Token = "xmHI9fjaLI8TGfCsCia-zr-bpPF3xSkop13GyJ5SVig=" and ((Pr.Name = "AllRights") or (Pr.Name = "get")) and ((Pr.tableName = "AllTables") or (Pr.tableName = "Managers"));

update Managers set roleId = "Rl_3" where Id = "Mn_2";