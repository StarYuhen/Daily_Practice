package CoreJavaVolumeOne;

// constructor test this
public class constructor {
    private String Name = "init";
    private int Age = 10;


    /**
     * 构造器的重载解析,同时分为静态，动态方法
     * 动态方法由VMT进行构造查找，每一个类都有一个这样的表，缩短查找时间，在动态方法中，父类的非静态方法可以被子类覆盖
     * 动态绑定，无须对现有代码进行修改就可以进行扩展。
     *
     */


    // init constructor
    // 构造函数没有返回类型，甚至不能使用 void 关键字来声明。当我们创建一个对象时，构造函数被调用并返回一个新的对象，这个对象是类的实例。


    // add code black
    {
        Age++;
    }

    // test differ constructor
    public constructor() {
        System.out.println("My is constructor()");
    }

    // 相同命名的构造器，传参值要不一样才行，返回类型不同的不行 ，重载匹配
//    public int DifferConstructor() {
//        return 1;
//    }

    public constructor(String string) {
        System.out.println("My is constructor string " + string);
    }

    public constructor(int age) {
        System.out.println("My is constructor int " + age);
    }

    public constructor(String name, int age) {
        System.out.println(name + age);
    }

    // test constructor this
    public constructor(double age) {
        this("Test ", 10);
    }

    public int getAge() {
        return Age;
    }

    public String getName() {
        return Name;
    }
}
