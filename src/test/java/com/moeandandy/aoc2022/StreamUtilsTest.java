package com.moeandandy.aoc2022;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.List;
import java.util.stream.Stream;
import org.junit.jupiter.api.Test;

class StreamUtilsTest {

    @Test
    void testSplitByBlanks() {
        List<List<String>> groups = StreamUtils.splitByBlanks(Stream.of("1", "2", "3", "", "4", "5", "6", "7"));
        assertEquals(2, groups.size());
        assertEquals(3, groups.get(0).size());
        assertEquals(4, groups.get(1).size());
    }
}