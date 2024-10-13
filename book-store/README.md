# Book-Store

Table :
一个书城网站可以设计以下几个主要表及其字段：

## 1. 用户表 (Users)

- **user_id** (主键)
- **username** (唯一)
- **password_hash**
- **email** (唯一)
- **role** (管理员、普通用户)
- **created_at**
- **updated_at**

### 2. 商品表 (Books)

- **book_id** (主键)
- **title**
- **author**
- **price**
- **stock**
- **category_id** (外键)
- **description**
- **cover_image**
- **created_at**
- **updated_at**

### 3. 分类表 (Categories)

- **category_id** (主键)
- **category_name**
- **created_at**
- **updated_at**

### 4. 订单表 (Orders)

- **order_id** (主键)
- **user_id** (外键)
- **order_status** (待发货、已发货、已完成)
- **total_amount**
- **created_at**
- **updated_at**

### 5. 订单详情表 (Order_Items)

- **order_item_id** (主键)
- **order_id** (外键)
- **book_id** (外键)
- **quantity**
- **price** (购买时的单价)

### 6. 评论表 (Reviews)

- **review_id** (主键)
- **user_id** (外键)
- **book_id** (外键)
- **rating** (评分)
- **comment**
- **created_at**

### 7. 地址表 (Addresses)

- **address_id** (主键)
- **user_id** (外键)
- **recipient_name**
- **address_line**
- **city**
- **postal_code**
- **phone_number**
- **created_at**

这样设计的数据库表结构可以覆盖基本的功能需求，具体字段可以根据实际情况进行调整或扩展。你有其他方面的需求吗？
