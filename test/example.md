# 服务器同步git更新

## 新建文件夹

`mkdir /repo/cst`
`cd /repo/cst`

## 初始化空git项目

`git init --bare`

### 在客户端添加远程项目

```bash
# user@local>
git init
git remote add origin root@server:/repo/cst/
git add .
git commit -m "init"
git push -u origin master
```

## 在服务端部署目录克隆项目

```bash
cd /var/www/html/deploy
git clone /repo/cst/
```


# 在服务端设置hooks

`vim /repo/cst/hooks/post-update`

在后面加入相应的更新代码:

```bash
cd /var/www/html/deploy
git pull origin master
```

## 在服务端部署目

```bash
cd /var/www/html/deploy
git clone /repo/cst/
```

### 在目

```bash
cd /var/www/html/deploy
git clone /repo/cst/
```
