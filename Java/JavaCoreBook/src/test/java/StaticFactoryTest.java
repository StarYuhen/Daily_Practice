import Book.CoreJavaVolumeOne.EffectiveJava.StaticFactory;
import CoreJavaVolumeOne.Factory;
import org.junit.Test;

public class StaticFactoryTest {
    // 测试静态工厂方法
    @Test
    public void StaticTest() throws Exception {
        Factory factory = StaticFactory.CreateStaticFactory(Factory.class);
        System.out.println(factory.getClass());
    }

    // 使用Builder方法
    @Test
    public void BuilderTest() {
        StaticFactory staticFactory = new StaticFactory
                .Builder("StarYuhen", 22)
                .setWeight(11.11)
                .setNation(StaticFactory.StaticFactoryEnum.Chine)
                .builder();
        System.out.println(staticFactory.getWeight());
    }
}
