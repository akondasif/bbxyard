package com.bbxyard.mp.auto_fill_meta.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

@Data
@TableName("mp_afm_user")
public class User {

    @TableId
    private Long id;

    @TableField("name")
    private String name;        // 值 映射

    private Integer age;

    private String email;

    @TableLogic
    private Integer deleted;    // 逻辑删除

    @Version
    private Integer version;    // 乐观锁 -> 版本

    private String operator;    // 操作员
}
