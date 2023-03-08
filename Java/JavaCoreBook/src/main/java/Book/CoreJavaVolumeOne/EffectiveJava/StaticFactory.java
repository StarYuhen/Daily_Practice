package Book.CoreJavaVolumeOne.EffectiveJava;


// 静态工厂方法，并取代构造器
public class StaticFactory {

    /*
     为什么使用静态工厂方法，而不是构造器
     1. 静态方法有自己的命名，不需要使用构造器进行签名，使用不同传参标识不同构造器的方法不利于她人使用
     2. 使用静态工厂的话，就不需要每次调用都创建一个方法了，可以减少消耗
     3. 可以利用反射返回任何子类型的对象
     4. 可以根据输入参数的不同而不同
     5. 编写包含该方法的类，返回对象类不用存在
     6. 限制部分方法被子类化
     */
    // 同时测试Builder模式
    /*
     为什么不使用每一个参数都set和get的方法
        这种模式被成为JavaBeans模式，有着严重的缺陷，构造方法在多次调用中被分割，所以构造过程就可能处于不一致的状态
        JavaBeans模式排除了让类不可变的可能性，而且需要保持线程安全
     */
    public enum StaticFactoryEnum {
        Chine, English, Germany, Japan
    }

    private String Name; // 姓名
    private int Age; // 年龄
    private double weight; // 重量
    private boolean gender; // 性别
    private String Race; // 种族
    private StaticFactoryEnum Nation; // 国家

    // 私有构造器

    private StaticFactory(Builder builder) {
        Name = builder.Name;
        Age = builder.Age;
        weight = builder.weight;
        gender = builder.gender;
        Race = builder.Race;
        Nation = builder.Nation;
    }

    // 编写的静态工厂方法
    // 使用泛型静态工厂返回任意传入类
    public static <T> T CreateStaticFactory(Class<T> tClass) throws Exception {
        // 注意这个反射方法被废弃了，所以换成最新的
//        return tClass.newInstance();
        return tClass.getDeclaredConstructor().newInstance();
    }


    // 用于测试数值


    public double getWeight() {
        return weight;
    }

    // 实现Builder模式
    public static class Builder {
        // 这里用于保存默认的配置，是否赋值不重要
        private String Name; // 姓名
        private int Age; // 年龄
        private double weight = 0.0; // 重量
        private boolean gender = false; // 性别
        private String Race = "未知"; // 种族
        private StaticFactoryEnum Nation = StaticFactoryEnum.Chine; // 国家

        // 实现构造函数
        public Builder(String Name, int Age) {
            this.Name = Name;
            this.Age = Age;
        }

        // 实现set方法
        public Builder setWeight(double weight) {
            this.weight = weight;
            return this;
        }

        public Builder setGender(boolean gender) {
            this.gender = gender;
            return this;
        }

        public Builder setRace(String race) {
            this.Race = race;
            return this;
        }

        public Builder setNation(StaticFactoryEnum Nation) {
            this.Nation = Nation;
            return this;
        }

        // 上级类
        public StaticFactory builder() {
            return new StaticFactory(this);
        }

    }


}
