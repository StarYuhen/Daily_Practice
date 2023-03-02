package yuhen.JavaCore;

// implements Cloneable in interface Or class
// deep copy function
public class DeepCopy implements Cloneable {
    // test deep or shallow copy
    private static final String Name = "Deep";
    private static int Age = 10;

    // deep copy function is clone
    @Override
    public DeepCopy clone() throws CloneNotSupportedException {
        return (DeepCopy) super.clone();
    }

    public static int getAge() {
        return Age;
    }

    public static void setAge(int age) {
        Age = age;
    }

}



