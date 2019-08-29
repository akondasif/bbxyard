package com.imooc.service.impl;

import com.github.pagehelper.PageHelper;
import com.imooc.mapper.SysUserMapper;
import com.imooc.mapper.SysUserMapperCustom;
import com.imooc.pojo.SysUser;
import com.imooc.service.SysUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.thymeleaf.util.StringUtils;
import tk.mybatis.mapper.entity.Example;

import java.util.List;

@Service
public class SysUserServiceImpl implements SysUserService {

    @Autowired
    private SysUserMapper sysUserMapper;

    @Autowired
    private SysUserMapperCustom sysUserMapperCustom;

    @Override
    public void saveUser(SysUser user) throws Exception {
        sysUserMapper.insert(user);
    }

    @Override
    public void updateUser(SysUser user) {
        sysUserMapper.updateByPrimaryKeySelective(user);
    }

    @Override
    public void deleteUser(String userId) {
        sysUserMapper.deleteByPrimaryKey(userId);
    }

    @Override
    public SysUser queryUserById(String userId) {
        return sysUserMapper.selectByPrimaryKey(userId);
    }

    @Override
    public List<SysUser> queryUserList(SysUser user) {
        Example example = new Example(SysUser.class);
        Example.Criteria criteria = example.createCriteria();

        if (!StringUtils.isEmptyOrWhitespace(user.getUsername())) {
            criteria.andLike("username", "%" + user.getUsername() + "%");
        }

        if (!StringUtils.isEmptyOrWhitespace(user.getNickname())) {
            criteria.andLike("nickname", "%" + user.getNickname() + "%");
        }

        // 排序
        example.orderBy("registTime").desc();
        List<SysUser> userList = sysUserMapper.selectByExample(example);
        return userList;
    }

    @Override
    public List<SysUser> queryUserListPaged(SysUser user, Integer page, Integer pageSize) {
        // 开始分页
        PageHelper.startPage(page, pageSize);

        Example example = new Example(SysUser.class);
        Example.Criteria criteria = example.createCriteria();

        if (!StringUtils.isEmptyOrWhitespace(user.getNickname())) {
            criteria.andLike("nickname", "%" + user.getNickname() + "%");
        }

        if (!StringUtils.isEmptyOrWhitespace(user.getUsername())) {
            criteria.andLike("username", "%" + user.getUsername() + "%");
        }

        // 排序
        example.orderBy("registTime").asc();
        List<SysUser> userList = sysUserMapper.selectByExample(example);
        return userList;
    }

    @Override
    public SysUser queryUserByIdCustom(String id) {
        List<SysUser> userList = sysUserMapperCustom.queryUserById(id);
        if (userList != null && !userList.isEmpty()) {
            return ((SysUser) userList.get(0));
        }
        return null;
    }

    @Override
    public void saveUserTransactional(SysUser user) {
        sysUserMapper.insert(user);
        int a = 1 / 0;
        user.setIsDelete(1);
        sysUserMapper.updateByPrimaryKeySelective(user);
    }
}
