package com.example;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

public class GreetingTest {

    @Test
    public void testSayHi() {
        assertEquals("Hi!", Greeting.sayHi());
    }
}
