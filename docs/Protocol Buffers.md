# Protocol Buffers

## 简介

> 官网 [protobuf](https://github.com/protocolbuffers/protobuf)
> 
> 下载地址 [protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
> 
> 相关文档：[programming-guides/proto3](https://protobuf.dev/programming-guides/proto3/)

Protocol Buffers 是一种与语言无关，平台无关的可扩展机制，用于序列化结构化数据。使用Protocol Buffers 可以一次定义结构化的数据，然后可以使用特殊生成的源代码轻松地在各种数据流中使用各种语言编写和读取结构化数据。
现在有许多框架等在使用Protocol Buffers。流行的RPC框架gRPC也是基于Protocol Buffers。
Protocol Buffers 目前有2和3两个版本号。

**优点**

* 性能较xml,json, thirft等好，效率高
* 代码生成机制，数据解析类自动生成
* 支持向后兼容和向前兼容
* 支持多种编程语言（java，c++，python）

**缺点**

* 二进制格式导致可读性差
* 缺乏自描述


## 文档结构
1. Protocol Buffers 使用 .proto 文件来保存文档。
2. Protocol Buffers文档的第一个非注释行，为版本申明，不填写的话默认为版本2。
```protobuf
syntax = "proto3";
/**或者**/
syntax = "proto2";

```
3. Packages 与Java的Package申明类型，Protocol Buffers 也可以声明package，来防止命名冲突。 不同的是 Packages是可选的。
```protobuf
package foo.bar;
message Open { 
  int32 open = 1;
}

```
使用的时候，也要加上命名空间
```protobuf
message Foo {
  foo.bar.Open open = 1;
}
```
在不同语言中生成的代码也是不一样的，java中为包名，c++为命名空间
4. 添加注释
和C/C++/JAVA注释语法一样
```protobuf
// 单行注释
/**
多行注释
多行注释
多行注释
**/

```
5. 接下来可以定义各种 消息和服务。
> 消息类似结构体、类等，服务是用来定义RPC通信的方法。

## 数据类型

Protocol Buffers 中数据类型分为标量类型和复合类型。类似，其它语言有基础类型和类等。复合类型，包括枚举和其他消息类型。

### 标量类型
标量消息字段可以具有以下类型之一 - 该表显示.proto文件中指定的类型，以及自动生成的类中的相应类型：


| .proto Type | Notes                                                        | C++ Type | Java/Kotlin Type[1] | Python Type[3]                  | Go Type | Ruby Type                   | C# Type    | PHP Type          | Dart Type |
| ----------- | ------------------------------------------------------------ | -------- | ------------------- | ------------------------------- | ------- | --------------------------- | ---------- | ----------------- | --------- |
| double      |                                                              | double   | double              | float                           | float64 | Float                       | double     | float             | double    |
| float       |                                                              | float    | float               | float                           | float32 | Float                       | float      | float             | double    |
| int32       | 使用可变长度编码。编码负数效率低下 数字 – 如果您的字段可能具有负值，请使用 sint32 相反 | int32    | int                 | int                             | int32   | Fixnum or Bignum (根据需要) | int        | integer           | int       |
| int64       | 使用可变长度编码。编码负数效率低下 数字 – 如果您的字段可能具有负值，请使用 sint64 相反 | int64    | long                | int/long[4]                     | int64   | Bignum                      | long       | integer/string[6] | Int64     |
| uint32      | 使用可变长度编码                                             | uint32   | int[2]              | int/long[4]                     | uint32  | Fixnum or Bignum (根据需要) | uint       | integer           | int       |
| uint64      | 使用可变长度编码                                             | uint64   | long[2]             | int/long[4]                     | uint64  | Bignum                      | ulong      | integer/string[6] | Int64     |
| sint32      | 使用可变长度编码。有符号的 int 值。这些 更多 比常规 int32 更有效地编码负数 | int32    | int                 | int                             | int32   | Fixnum or Bignum (根据需要) | int        | integer           | int       |
| sint64      | 使用可变长度编码。有符号的 int 值。这些 更多 比常规 int64 更有效地编码负数 | int64    | long                | int/long[4]                     | int64   | Bignum                      | long       | integer/string[6] | Int64     |
| fixed32     | 始终为四个字节。比 uint32 更有效，如果值经常 大于 228        | uint32   | int[2]              | int/long[4]                     | uint32  | Fixnum or Bignum (根据需要) | uint       | integer           | int       |
| fixed64     | 始终为 8 个字节。如果值经常 大于 2^56                        | uint64   | long[2]             | int/long[4]                     | uint64  | Bignum                      | ulong      | integer/string[6] | Int64     |
| sfixed32    | 始终为四个字节                                               | int32    | int                 | int                             | int32   | Fixnum or Bignum (根据需要) | int        | integer           | int       |
| sfixed64    | 始终为 8 个字节。                                            | int64    | long                | int/long[4]                     | int64   | Bignum                      | long       | integer/string[6] | Int64     |
| bool        |                                                              | bool     | boolean             | bool                            | bool    | TrueClass/FalseClass        | bool       | boolean           | bool      |
| string      | 字符串必须始终包含 UTF-8 编码或 7 位 ASCII 文本，并且不能 大于 2^32 | string   | String              | str/unicode[5]                  | string  | String (UTF-8)              | string     | string            | String    |
| bytes       | 可以包含任何长度不超过 2 的任意字节序列32.                   | string   | ByteString          | str (Python 2) bytes (Python 3) | []byte  | String (ASCII-8BIT)         | ByteString | string            | List      |

> 更多说明查看[proto3](https://protobuf.dev/programming-guides/proto3/)

### 复合类型

1. 枚举
在 Proto Buffers 中，我们可以定义枚举和枚举类型，
```protobuf
enum Corpus {
  UNIVERSAL = 0;
  WEB = 1;
  IMAGES = 2;
  LOCAL = 3;
  NEWS = 4;
  PRODUCTS = 5;
  VIDEO = 6;
}

Corpus corpus = 4;
```
枚举定义在一个消息内部或消息外部都是可以的，如果枚举是 定义在 message 内部，而其他 message 又想使用，那么可以通过 MessageType.EnumType 的方式引用。

定义枚举的时候，我们要保证第一个枚举值必须是0，枚举值不能重复，除非使用 option allow_alias = true 选项来开启别名。
```protobuf
enum EnumAllowingAlias {
    option allow_alias = true;
    UNKNOWN = 0;
    STARTED = 1;
    RUNNING = 1;
}
```

枚举值的范围是32-bit integer，但因为枚举值使用变长编码，所以不推荐使用负数作为枚举值，因为这会带来效率问题

## 定义消息

Proto Buffers 使用message定义消息。例如:
```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

```
* 该文件的第一行指定使用proto3语法：如果不填，Protocol Buffers编译器将默认使用proto2。这必须是文件的第一个非空的非注释行。
* 所述SearchRequest消息定义指定了三个字段（名称/值对），一个查询字符串，结果页面数以及每页的结果个数。每个字段都有一个名称和类型和编号。

### 字段编号
如您所见，消息定义中的每个字段都有唯一的编号。这些字段编号用于以消息二进制格式标识字段，并且在使用消息类型后不应更改。 请注意，1到15范围内的字段编号需要一个字节进行编码，包括字段编号和字段类型（您可以在协议缓冲区编码中找到更多相关信息）。16到2047范围内的字段编号占用两个字节。因此，您应该为非常频繁出现的消息元素保留数字1到15。请记住为将来可能添加的常用元素留出一些空间。

您可以指定的最小字段数为1，最大字段数为 2 29 - 1或536,870,911。您也不能使用数字19000到19999 （FieldDescriptor::kFirstReservedNumber ~ FieldDescriptor::kLastReservedNumber），因为它们是为Protocol Buffers实现保留的 - 如果您使用其中一个保留号码，Protocol Buffers编译器会报错。同样，您不能使用任何以前保留的字段编号。

### 指定字段规则
消息字段可以是以下之一：

1. singular：格式良好的消息可以包含该字段中的零个或一个（但不超过一个）。
2. repeated：此字段可以在格式良好的消息中重复任意次数（包括零）。将保留重复值的顺序。

在proto3中，repeated标量数字类型的字段为默认使用编码。

在proto2中:

* required：必须有一个
* optional：0或者1个
* repeated：任意数量（包括0）


### 保留字段

保留变量不被使用

如果通过完全删除字段或将其注释来更新消息类型，则未来用户可以在对类型进行自己的更新时重用字段编号。如果以后加载相同的旧版本，这可能会导致严重问题，包括数据损坏，隐私错误等。确保不会发生这种情况的一种方法是指定已删除字段的字段编号（和/或名称，这也可能导致JSON序列化问题）reserved。如果将来的任何用户尝试使用这些字段标识符，协议缓冲编译器将会报错。

```protobuf
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}

```

### 默认值

解析消息时，如果编码消息不包含特定的单数元素，则解析对象中的相应字段将设置为该字段的默认值。这些默认值是特定于类型的：

* 对于字符串，默认值为空字符串。
* 对于字节，默认值为空字节。
* 对于bools，默认值为false。
* 对于数字类型，默认值为零。
* 对于枚举，默认值是第一个定义的枚举值，该值必须为0。
* 对于消息字段，未设置该字段。它的确切值取决于语言。有关详细信息， 请参阅生成的代码指
* 重复字段的默认值为空（通常是相应语言的空列表）。

请注意，对于标量消息字段，一旦解析了消息，就无法确定字段是否显式设置为默认值（例如比如bool，是否设置了布尔值false）或者根本没有设置：您应该记住这一点在定义消息类型时。
导入定义
默认情况下，您只能使用直接导入.proto文件中的定义。但是，有时您可能需要将.proto文件移动到新位置。.proto现在，您可以.proto在旧位置放置一个虚拟文件，以使用该import public概念将所有导入转发到新位置，而不是直接移动文件并在一次更改中更新所有调用站点。**import public任何导入包含该import public语句的proto的人都可以传递依赖关系。** 

例如：

```protobuf
// new.proto
//所有定义都移到这里



// old.proto
//这是所有客户端都要导入的原型。
import public "new.proto";
import "other.proto";


// client.proto
import "old.proto";

//您使用old.proto和new.proto中的定义，但不使用other.proto
```

### 嵌套类型

你可以在其他消息类型中定义、使用消息类型，在下面的例子中，Result消息就定义在SearchResponse消息内，如：

```protobuf
message SearchResponse {
  
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```
如果要在其父消息类型之外重用此消息类型，请将其称为： Parent.Type


### 使用proto2消息类型

可以导入proto2消息类型并在proto3消息中使用它们，反之亦然。但是，proto2枚举不能直接用于proto3语法（如果导入的proto2消息使用它们就可以了）。

### 更新消息类型

如果现有的消息类型不再满足您的所有需求 - 例如，您希望消息格式具有额外的字段 - 但您仍然希望使用使用旧格式创建的代码，请不要担心！在不破坏任何现有代码的情况下更新消息类型非常简单。请记住以下规则：

* 请勿更改任何现有字段的字段编号。
* 如果添加新字段，则使用“旧”消息格式按代码序列化的任何消息仍可由新生成的代码进行解析。您应该记住这些元素的默认值，以便新代码可以正确地与旧代码生成的消息进行交互。同样，您的新代码创建的消息可以由旧代码解析：旧的二进制文件在解析时只是忽略新字段。有关详细信息，请参阅“ 未知字段”部分
* 只要在更新的消息类型中不再使用字段编号，就可以删除字段。您可能希望重命名该字段，可能添加前缀“OBSOLETE_”，或者保留字段编号，以便您的未来用户.proto不会意外地重复使用该编号。
* int32，uint32，int64，uint64，和bool都是兼容的-这意味着你可以改变这些类型到另一个的一个场不破坏forwards-或向后兼容。如果从导线中解析出一个不符合相应类型的数字，您将获得与在C ++中将该数字转换为该类型相同的效果（例如，如果将64位数字作为int32读取，它将被截断为32位）。
* sint32并且sint64彼此兼容但与其他整数类型不兼容。
* stringbytes只要字节是有效的UTF-8 ，它们是兼容的。
* bytes如果字节包含消息的编码版本，则嵌入消息是兼容的。
* fixed32与兼容sfixed32，并fixed64用sfixed64。
* enum与兼容int32，uint32，int64，和uint64电线格式条款（注意，如果他们不适合的值将被截断）。但请注意，在反序列化消息时，客户端代码可能会以不同方式对待它们：例如，enum将在消息中保留未识别的proto3 类型，但在反序列化消息时如何表示这种类型取决于语言。Int字段总是保留它们的价值。
* 将单个值更改为新 成员oneof是安全且二进制兼容的。oneof如果您确定没有代码一次设置多个字段，则将多个字段移动到新字段可能是安全的。将任何字段移动到现有字段oneof并不安全。

### 未知字段

未知字段是格式良好的协议缓冲区序列化数据，表示解析器无法识别的字段。例如，当旧二进制文件解析具有新字段的新二进制文件发送的数据时，这些新字段将成为旧二进制文件中的未知字段。

Proto3实现可以成功解析具有未知字段的消息，但是，实现可能支持也可能不支持保留那些未知字段。您不应该依赖保留或删除的未知字段。对于大多数Google协议缓冲区实现，未知字段在proto3中无法通过相应的proto运行时访问，并且在反序列化时被删除和遗忘。这与proto2的行为不同，其中未知字段始终与消息一起保留和序列化。

**Any**

**Oneof**

如果您有一个包含许多字段的消息，并且最多只能同时设置一个字段，则可以使用oneof功能强制执行此行为并节省内存。

除了一个共享内存中的所有字段之外，其中一个字段类似于常规字段，并且最多可以同时设置一个字段。设置oneof的任何成员会自动清除所有其他成员。您可以使用特殊case()或WhichOneof()方法检查oneof中的哪个值（如果有），具体取决于您选择的语言。
使用Oneof

要在您中定义oneof，请.proto使用oneof关键字后跟您的oneof名称，在这种情况下test_oneof:


**map**

如果要在数据定义中创建关联映射，协议缓冲区提供了一种方便的快捷方式语法：
```protobuf
map < key_type ，value_type > map_field = N ;
```
...其中key_type可以是任何整数或字符串类型（因此，除了浮点类型之外的任何标量类型bytes）。请注意，枚举不是有效的key_type。的value_type可以是任何类型的除另一地图。
因此，例如，如果要创建项目映射，其中每条Project消息都与字符串键相关联，则可以像下面这样定义它：

```protobuf
map < string ，Project > projects = 3 ;
```

地图字段不能repeated。地图值的有线格式排序和地图迭代排序未定义，因此您不能依赖于特定顺序的地图项目。为a生成文本格式时.proto，地图按键排序。数字键按数字排序。从线路解析或合并时，如果有重复的映射键，则使用最后看到的键。从文本格式解析映射时，如果存在重复键，则解析可能会失败。如果为映射字段提供键但没有值，则字段序列化时的行为取决于语言。在C ++，Java和Python中，类型的默认值是序列化的，而在其他语言中没有任何序列化。生成的地图API目前可用于所有proto3支持的语言。您可以在相关API参考中找到有关所选语言的地图API的更多信息。

## 定义服务

Proto Buffers 可以使用在RPC（远程过程调用）环境中。 要定义RPC方法，需要使用service来定义服务，并在里面定义方法。

Proto Buffers 编译器会根据选择的语言生成服务接口代码和存根（Stub）.

```protobuf
service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}
```
Proto Buffers 可以直接在gRPC框架使用，当然，也可以自定义RPC实现，不使用gRPC。

## JSON映射

Proto Buffers 3 增加了 Json 映射，可实现protobuf与json互相转换。

## proto3 与 proto2 的区别

* 在第一行非空白非注释行，必须写：syntax = "proto3";
* 移除了 “required”，并把 “optional” 改名为"singular"；
* 语言增加 Go、Ruby、JavaNano 支持；
* 移除了 default 选项；
* 枚举类型的第一个字段必须为 0 ；
* 移除了对分组的支持；
* 分组的功能完全可以用消息嵌套的方式来实现，
* 旧代码在解析新增字段时，会把不认识的字段丢弃，再序列化后新增的字段就没了；
* 移除了对扩展的支持，新增了 Any 类型；
* Any 类型是用来替代 proto2 中的扩展。
* 增加了 JSON 映射特性；

