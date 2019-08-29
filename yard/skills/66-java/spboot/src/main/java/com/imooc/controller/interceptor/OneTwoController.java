package com.imooc.controller.interceptor;

import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("hook/one-two")
public class OneTwoController {

    @RequestMapping("index")
    public String index(ModelMap map) {
        map.addAttribute("name", "hook/one-two/index");
        return "thymeleaf/index";
    }
}
