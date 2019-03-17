<?php
/**
 * Created by PhpStorm.
 * User: bbxyard
 * Date: 19-3-7
 * Time: 下午2:17
 */

class StrTest extends \PHPUnit\Framework\TestCase
{
    public function testStrReplace() {
        $s1 = ":a:b:c:";
        $s2 = str_replace(':', '/', $s1);
        $this->assertEquals('/a/b/c/', $s2);
    }

    public function testEncode()
    {
        // md5
        print('md5(): ' . md5('') . "\n");
        print_r('md5(1): ' . md5('1') . "\n");
        printf("md5(123456): %s\n", md5('123456'));
    }

    public function testNumberConv()
    {
        $farr = [];
        $farr[] = floatval('3.14');
        $farr[] = floatval(3.14);
        array_push($farr, floatval(3));
        var_dump($farr);
    }

    public function testF()
    {
        $file = '/etc/passwd';
        echo @file_get_contents($file);
        printf("ctime: %s, mtime=%s, atime=%s\n", filectime($file), filemtime($file), fileatime($file));
    }

    public function testMisc()
    {
        $aa = 60;
        $bb = time();
        $cc = uniqid('hay-');
        $arr = compact('aa', 'bb', 'cc');
        // list($x, $y, $z) = $arr;
        list($x, , $y) = ['boxu', '24', '1985'];
        var_dump($arr);
        echo "aa = $aa - $x, bb = $bb - $y\n";
        echo 'hello '.'xxx'.$aa."\n";
        echo date('Y-m-d H:i:s',time())."\n";
        echo substr(md5('hallo welt'), 12, 8)."\n";
        echo str_replace('Das ist', 'This is', '##Das ist China##')."\n";
        echo str_replace('Das2 ist', 'This is', '##Das ist China##')."\n";
        var_dump(strstr('hallo', 'abc'));
        var_dump(strpos('hallo', 'hall'));
        // var_dump(compact('a', 'b', 'c'));
        echo uniqid('#hallo welt#')."\n";
        echo uniqid()."\n";
        echo uniqid(true)."\n";
        echo uniqid(1)."\n";
        echo "PHP版本:".phpversion();
    }
}