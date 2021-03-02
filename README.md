# Sample-Go

<br>

*auto created by [readme_catalog_builder](src/main/readme_catalog_builder.go)*

<br>

### Go 基本使用示例代码
#### Go 数据类型、数据结构
* [datatype_spl.go](./src/spl.assad/main/a1_datatype_spl.go): go 基本数据类型
* [datatype_num_spl.go](./src/spl.assad/main/a2_datatype_num_spl.go): 数值类型使用
* [datatype_string_spl.go](./src/spl.assad/main/a3_datatype_string_spl.go): 字符串类型使用
* [datetype_datetime_spl.go](./src/spl.assad/main/a4_datetype_datetime_spl.go): 时间日期类型使用
* [datatype_array_slice_spl.go](./src/spl.assad/main/a5_datatype_array_slice_spl.go): 数组，切片（slice）使用示例
* [datatype_map_spl.go](./src/spl.assad/main/a6_datatype_map_spl.go): Map 使用示例
* [datatype_container_list_range_spl.go](./src/spl.assad/main/a7_datatype_container_list_range_spl.go): container 标准包提供的 list（双链表）、ring（环形链表） 集合操作
#### Go 控制结构语法
* [control_structure_spl.go](./src/spl.assad/main/b1_control_structure_spl.go): 控制结构语句（if-else, for-i, for-range, for-condition, goto）
#### Go 函数
* [func_call_spl.go](./src/spl.assad/main/c1_func_call_spl.go): 函数调用示例
* [func_feature_spl.go](./src/spl.assad/main/c2_func_feature_spl.go): 函数特性示例
* [func_callback_spl.go](./src/spl.assad/main/c3_func_callback_spl.go): lambda 回调函数示例（使用其他函数作为入参）
* [func_closure_spl.go](./src/spl.assad/main/c4_func_closure_spl.go): closure 闭包（匿名函数）示例
* [func_debug_spl.go](./src/spl.assad/main/c5_func_debug_spl.go): 函数调试, 打印函数执行位置，统计函数执行时间
* [func_fibonacci_spl.go](./src/spl.assad/main/c6_func_fibonacci_spl.go): 递归函数示例， fibonacci 计算
* [func_fibonacci_memoization_spl.go](./src/spl.assad/main/c7_func_fibonacci_memoization_spl.go): 递归函数示例， fibonacci 计算，使用内存缓存
#### Go 标准包
* [standpkg_os_cmd_spl.go](./src/spl.assad/main/d1_standpkg_os_cmd_spl.go): go os 包调用，运行系统命令等示例
* [standpkg_cmd_read_spl.go](./src/spl.assad/main/d2_standpkg_cmd_read_spl.go): 从控制台读取输入（os.Stdin）
* [standpkg_cmd_arg_spl.go](./src/spl.assad/main/d3_standpkg_cmd_arg_spl.go): 从启动命令行读取参数 os.Args
* [standpkg_regexp_spl.go](./src/spl.assad/main/d4_standpkg_regexp_spl.go): go regexp 正则匹配
* [standpkg_math_big_spl.go](./src/spl.assad/main/d5_standpkg_math_big_spl.go): go math/big 精密计算
* [standpkg_json_spl.go](./src/spl.assad/main/d6_standpkg_json_spl.go): go json 处理（encoding/json）
* [standpkg_xml_spl.go](./src/spl.assad/main/d7_standpkg_xml_spl.go): go xml 处理
* [standpkg_encode_decode_spl.go](./src/spl.assad/main/d8_standpkg_encode_decode_spl.go): go 加解密处理（sha、md5 等）
* [thirdparty_vandor_spl.go](./src/spl.assad/main/d9_thirdparty_vandor_spl.go): vendor 安装、调用第三方包依赖
#### Go 面向对象
* [oop_struct_method_spl.go](./src/spl.assad/main/f1_oop_struct_method_spl.go): go struct 结构体示例
* [oop_struct_Getter_Setter.go](./src/spl.assad/main/f2_oop_struct_Getter_Setter.go): go struct 结构体示例（字段 Getter、Setter）
* [oop_struct_embedded_spl.go](./src/spl.assad/main/f3_oop_struct_embedded_spl.go): go struct 内嵌结构示例
* [oop_struct_anonymous_field_spl.go](./src/spl.assad/main/f4_oop_struct_anonymous_field_spl.go): go struct 匿名字段、匿名内嵌结构体 示例
* [oop_struct_tag.go](./src/spl.assad/main/f5_oop_struct_tag.go): go struct 标签定义、反射获取示例
* [oop_runtime_gc_spl.go](./src/spl.assad/main/f6_oop_runtime_gc_spl.go): go 显式调用 runtime gc
* [oop_interface_spl.go](./src/spl.assad/main/f7_oop_interface_spl.go): interface 接口使用示例
* [oop_interface_nest_spl.go](./src/spl.assad/main/f8_oop_interface_nest_spl.go): interface 接口嵌套示例
* [oop_interface_type_judge_spl.go](./src/spl.assad/main/f9_oop_interface_type_judge_spl.go): go 接口类型判断
* [oop_node_struct_spl.go](./src/spl.assad/main/f10_oop_node_struct_spl.go): 一个简单的 Node 节点数据结构
* [oop_reflect.go](./src/spl.assad/main/f11_oop_reflect.go): go 类型反射示例（TypeOf，ValueOf）
* [oop_gob_spl.go](./src/spl.assad/main/f12_oop_gob_spl.go): go gob 二进制序列化、反序列化
#### Go 文件处理
* [file_read_spl.go](./src/spl.assad/main/g1_file_read_spl.go): 文件读取示例
* [file_write_spl.go](./src/spl.assad/main/g2_file_write_spl.go): 文件写入示例
* [file_copy_spl.go](./src/spl.assad/main/g3_file_copy_spl.go): 文件拷贝示例
  <br>




<br>
<br>
