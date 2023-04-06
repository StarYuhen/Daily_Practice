package Book.CoreJavaVolumeOne.EffectiveJava;

import java.lang.ref.Cleaner;

// 使用finalizer和cleaner机制，同时使用Cleaner
public class NoFinalizer implements AutoCloseable {
    // 实现AutoCloseable 声明本地对等类
    private static final Cleaner cleaner = Cleaner.create();
    private final Cleaner.Cleanable cleanable;

    public NoFinalizer() {
        State state = new State(10);
        cleanable = cleaner.register(this, state);
    }

    @Override
    public void close() {
        // 通知Gc回收
        cleanable.clean();
    }

    // 启动一个任务
    private static class State implements Runnable {
        int Number = 0;

        State(int number) {
            this.Number = number;
        }

        @Override
        public void run() {
            System.out.println("Run Cleaning room");
            for (int i = 0; i < 100000; i++) {
                System.out.println(i);
            }
        }

    }
}
