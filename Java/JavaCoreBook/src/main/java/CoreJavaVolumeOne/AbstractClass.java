package CoreJavaVolumeOne;

// 实现抽象类
public abstract class AbstractClass {
    // 一个抽象方法 更像是一个占位角色
    public abstract String getAbstractName();
    // 测试抽象类的子类
    public static void main(String[] args) {
        TestName test = new TestName();
        System.out.println(test.getAbstractName());
    }
}

// 新建子类来继承他
class TestName extends AbstractClass {
    private static String name = "StarYuhen";
    // 实现父类的抽象方法

    @Override
    public String getAbstractName() {
        return name;
    }
}