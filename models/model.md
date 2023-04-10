# 集合列表

## 用户表
```json
{
    "accout": "账号",
    "password": "密码",
    "nickname": "昵称",
    "sex": "1",
    "email": "邮箱",
    "avartar": "头像",
    "created_at": "1",
    "updated_at": "1"
}
```

## 消息集合
```json
{
    "user_identity":"用户的唯一标识",
    "room_identity":"房间的唯一标识",
    "data":"发送的数据",
    "created_at":1,
    "updated_at":1
}
```

## 房间集合
```json
{
    "number":"房间号",
    "name":"房间名",
    "info":"房间简介",
    "user_identity":"房间创建者的唯一标识",
    "created_at":1,
    "updated_at":1
}
```

## 用户-房间-消息关联表
```json
{
    "user_identity":"用户的唯一标识",
    "room_identity":"房间的唯一标识",
    "message_identity":"消息的唯一标识",
    "created_at":1,
    "updated_at":1
}
```