package CoreJavaVolumeOne;

// implements Cloneable in interface Or class
// deep copy function
public class DeepCopy implements Cloneable {
    // test deep or shallow copy
    private static final String Name = "Deep";
    /*
    private static int Age = 10;
    1. 即使使用了深拷贝的clone的情况下，依然能够更改这个属性
        原因是因为： static静态资源属性，他会在申明的第一个时刻将其写入方法区，改变他就会改变所有对象的属性值
     */
    private int Age = 10;
    public int Weight = 10;

    // deep copy function is clone
    @Override
    public DeepCopy clone() throws CloneNotSupportedException {
        return (DeepCopy) super.clone();
    }

    public int getAge() {
        return Age;
    }

    public void setAge(int age) {
        this.Age = age;
    }

}



