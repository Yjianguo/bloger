-- MySQL dump 10.13  Distrib 8.0.21, for macos10.15 (x86_64)
--
-- Host: 10.227.10.78    Database: blog
-- ------------------------------------------------------
-- Server version	5.7.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) NOT NULL COMMENT '分类id',
  `content` longtext CHARACTER SET utf8mb4 NOT NULL COMMENT '文章内容',
  `title` varchar(1024) CHARACTER SET utf8mb4 NOT NULL COMMENT '文章标题',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `username` varchar(128) CHARACTER SET utf8mb4 NOT NULL COMMENT '作者',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
  `summary` varchar(256) CHARACTER SET utf8mb4 NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1 COMMENT='文章表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (1,3,'OKR（7-8月）\n1、持续推进业务多机房容灾体系建设（国内和海外）\n（1）容灾预案的管理、维护，预案有效性验证\n（2）核心服务的降级、跨机房业务的专项治理\n（3）故障演练、预案演练\n\n2、业务稳定性建设和保障\n（1）业务稳定性建设的规范、流程\n（2）抖火核心服务/接口的梳理和治理(最小可用)\n（3）核心场景监控报警治理、oncall\n\n3、社交中台、核心服务容灾预案建设？\n4、其他','OKR',100,0,'yuanjg',1,'okr... ...','2020-09-09 08:54:19',NULL),(2,1,'#### 基本操作\n\n1. **git init** 初始化仓库\n2. **git status**  查看仓库状态\n3. **git add** 向暂存区添加文件\n4. **git commit** 保存仓库中的历史记录\n\n```shell\nyuanjg@Bytedance % git commit -m \"First commit\"\n[master b58fb41] First commit\n 1 file changed, 0 insertions(+), 0 deletions(-)\n create mode 100644 README.md\n \n```\n\n5. **git log** 查看提交日志，可以通过 --help 查看更多参数\n\n```shell\nyuanjg@Bytedance % git log\ncommit b58fb41768458b787b7c94eb7da842127b1b8e82 (HEAD -> master)\nAuthor: yuanjianguo <yuanjianguo@bytedance.com>\nDate:   Fri Aug 28 10:40:45 2020 +0800\n\n    First commit\n    \n    Change-Id: Icad5f0fe8a2c5270758970a713cda386a128fd84\n\ncommit 6abc46fce8df9a19b2ddcfadff18e0ccba1e1ec0\nAuthor: yuanjianguo <yuanjianguo@bytedance.com>\nDate:   Thu Aug 27 21:34:17 2020 +0800\n\n    new projects\n    \n    Change-Id: I33e27925be7767ce0b00d29322ce527fbff4e4f8\n\n```\n\n6. **git diff** 查看更改前后的差别','git基本操作',10000,100,'wangxl',1,'git... ....','2020-09-09 09:01:14',NULL),(3,5,'ecosystem [ˈiːkəʊsɪstəm] 生态圈\nmentor	[ˈmentɔː(r)]  导师/顾问\nprogress [ˈprəʊɡres] 进步，进展\nego [\'i:ɡəu, \'eɡəu] 自我/自负/自我意识\ntutorial [tuːˈtɔːriəl]  教程，学习指南\ncompiler [kəmˈpaɪlər] 编译器/编译者\ngenerated [\'dʒɛnə,retɪd]  生成/产生/引发\nconsistency [kənˈsɪstənsi] 一致性；稠度；相容性\ncomponent  [kəmˈpəʊnənt] 组成部分、成分、组件、元件；组成的、构成的\nevaluate  [ɪˈvæljueɪt] 评估、评价、求...的值\n-- We need to evaluate how well the policy is working.\nexternal  [ɪkˈstɜːnl] 外部的、表面的；外部、外观\nalgorithm [ˈælɡərɪðəm] 算法、运算法则\ninspection  [ɪnˈspekʃn]  检查、视察\nconstraints [kən\'streint] 约束、限制，约束条件 \nexclude  [ɪkˈskluːd]  排除、排斥\nassignment  [əˈsaɪnmənt] 分配、任务、作业、功课\nretrieve  [rɪˈtriːv]   检索、恢复、取回 \n-- The program allows you to retrieve items quickly by searching under a keywords. \nplaceholder [ˈpleɪshoʊldər] 占位符\nconduct [kənˈdʌkt]  组织、实施\ncompetitive [kəmˈpetətɪv] 竞争的、比赛的\nelegant  [ˈelɪɡənt]  高雅的、讲究的、简略的\nShe was tall and elegant。她身材修长、优雅大方。\nmanpower 人力、人力资源、劳动力\nschedule  [ˈʃedjuːl] 安排、日程表、将...列入计划表或清单\nhectic  [ˈhektɪk] 紧张的、忙碌的。脸红、肺病的\nI have a hectic schedule for the next few days.\nofficially  [əˈfɪʃəli] 官方的、正式的\nparallelism [ˈpærəlelɪzəm] 平行、并行\nconcurrency  [kənˈkʌrənsɪ] 并发\nexpand  [ɪkˈspænd] 扩大 发展、 使膨胀\nstakeholder  [ˈsteɪkhəʊldə(r)] 利益相关者\nmarshal  [ˈmɑːʃl]   元帅、编排，排列，序列化 \nspy [spaɪ]  间谍，从事间谍活动\ninvolve [ɪnˈvɒlv] 牵涉，包含\nconvert [kənˈvɜːt] 转换，使转换\nduplicate [ˈdjuːplɪkeɪt] 复制的，复制品\nfragment   [\'frægmənt] 碎片、片段、不完整\nbare  [beə(r)] 空的，赤裸的。露出，使赤裸\nclause [klɔːz]  条款\nleggings  [ˈleɡɪŋz] 裹腿 紧身裤\nredundant   [rɪˈdʌndənt]  多余的，过剩的，被解雇的、失业的，冗长的、繁杂的。\nprocedure [prəˈsiːdʒə(r)] 程序、手续、步骤 ','单词本',1000,111,'xialu',1,'ecosystem ... ...','2020-09-09 09:03:01',NULL),(4,6,'KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS','html',1111,1111,'xiaowang',1,'kkkk','2020-09-09 09:59:22',NULL),(5,6,'KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS','html',1111,1111,'xiaowang',1,'kkkk','2020-09-10 08:02:50',NULL),(6,6,'KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS','html',1111,1111,'xiaowang',1,'kkkk','2020-09-15 07:21:15',NULL),(7,6,'KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS','html',1111,1111,'xiaowang',1,'kkkk','2020-09-15 07:22:56',NULL),(8,4,'<p><strong>常见的博客系统</strong></p><h3>&nbsp;</h3><p>Windows Server 2012 Standard标准版密钥：</p><p>[Key]：J7TJK-NQPGQ-Q7VRH-G3B93-2WCQD</p><p>[Key]：R7CV7-NWJYG-GYY64-MXBWC-D669Q</p><p>Mblog</p><p>Halo</p><p>oneblog</p><p><strong>Java系列</strong></p><p>Hugo go语言</p><p>Hexo node.js</p><p><strong>静态页面系列</strong></p><p>Typeecho</p><p>Z-blog</p><p>Wordpress</p><p><strong>PHP系列</strong></p>','飞蛾扑火的SB',0,0,'黄花大闺女',1,'p><strong>常见的博客系统</strong></p><h3>&nbsp;</h3><p>Windows Server 2012 Standard标准版','2020-09-15 12:47:12',NULL),(9,6,'KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS','html',1111,1111,'xiaowang',1,'kkkk','2020-09-15 14:06:55',NULL),(10,6,'KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS','html',1111,1111,'xiaowang',1,'kkkk','2020-09-15 14:09:22',NULL);
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '分类名字',
  `category_no` int(10) NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 COMMENT='分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'Python',100,'2020-09-08 14:11:03',NULL),(2,'Java',101,'2020-09-08 14:11:04',NULL),(3,'Go',102,'2020-09-08 14:11:04',NULL),(4,'Linux',103,'2020-09-08 14:11:04',NULL),(5,'English',104,'2020-09-09 09:03:34',NULL);
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text CHARACTER SET utf8mb4 NOT NULL COMMENT '评论内容',
  `username` varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态：0，删除：1，保留',
  `article_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1 COMMENT='评论表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,'666,666','xxx','2020-09-09 09:06:08',1,1),(2,'写的真不错，求更新。','花花世界','2020-09-09 09:06:08',1,2),(3,'终于有人说明白这事儿了，赞?','小霸王','2020-09-09 09:06:08',1,3),(4,'jjssdskjncksksdkskds','lnkkhj','2020-09-15 22:06:56',1,3),(5,'jjssdskjncksksdkskds','lnkkhj','2020-09-15 22:09:23',1,3),(6,'好厉害啊','好厉害','2020-09-15 22:31:50',1,0),(7,'快捷方式贷款福克斯的科技风科技的时刻','的技术打假赛的','2020-09-15 22:35:23',1,0),(8,'lfnafsdnfkjakjnkjbsfasjkfbbs  ','dijsaoodsldn','2020-09-15 22:37:16',1,0),(9,'dasdasdasda ','dasdadas','2020-09-15 22:39:18',1,0),(10,'asdasdasdasdasdadad','dasdadas','2020-09-15 22:42:44',1,0),(11,'第三代上档次','胜多负少的','2020-09-15 22:45:50',1,0),(12,'ddsddsdsdsdfdsfs','ssdsass','2020-09-15 22:46:34',1,0),(13,'lsdllsaljdlskjclksdjldjaskkd;sakl','yuanjgjjjjjjjjjj','2020-09-16 10:07:11',1,8);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `leaves`
--

DROP TABLE IF EXISTS `leaves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `leaves` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '留言id',
  `username` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '留言用户',
  `email` varchar(255) DEFAULT NULL COMMENT '用户邮箱',
  `content` text CHARACTER SET utf8mb4 COMMENT '留言内容',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '留言发布时间',
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1 COMMENT='留言表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `leaves`
--

LOCK TABLES `leaves` WRITE;
/*!40000 ALTER TABLE `leaves` DISABLE KEYS */;
INSERT INTO `leaves` VALUES (1,'xxx','xxx@163.com','666666','2020-09-09 09:09:10',NULL),(2,'小霸王','xbw@163.com','大佬，求友链。','2020-09-09 09:09:10',NULL),(3,'yuanjg','yuanjg@163.com','写的超级棒，nice，学习了','2020-09-15 15:21:16',NULL),(4,'yuanjg','yuanjg@163.com','写的超级棒，nice，学习了','2020-09-15 15:22:57',NULL),(5,'喜爱你 ','xihai@163.com','拉你了三年内会计师那可就从哪开始的存款时可参考时刻表','2020-09-15 15:52:22',NULL),(6,'秀安安','xiuanan@163.com','你说的内陆费年四季度科技城看到你','2020-09-15 15:56:55',NULL),(7,'sjdkjsdjksjdbsbjdbsjbfjbjfs','sdsnkcnsjkknkskk@163.com','fjsdnfkkfksnkdn s ck  sdkcksdkckdsabjjsc,b','2020-09-15 15:58:10',NULL),(8,'fsfsfsfadfs','ddlsadfslfnlsdnl@12222.com','dsspojofjskfnkknfnjansfdd','2020-09-15 15:59:04',NULL),(9,'oijpohbj knno','njijkkjnkjnkjnk@111.com','dsfsfsafsadfafas','2020-09-15 15:59:24',NULL),(10,'yuanjg','yuanjg@163.com','写的超级棒，nice，学习了','2020-09-15 22:06:56',NULL),(11,'yuanjg','yuanjg@163.com','写的超级棒，nice，学习了','2020-09-15 22:09:23',NULL);
/*!40000 ALTER TABLE `leaves` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-09-16 10:54:40
