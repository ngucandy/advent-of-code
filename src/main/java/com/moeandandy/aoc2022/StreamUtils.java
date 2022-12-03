package com.moeandandy.aoc2022;

import java.util.LinkedList;
import java.util.List;
import java.util.stream.Stream;

public class StreamUtils {

    public static List<List<String>> splitByBlanks(Stream<String> stringStream) {
        List<List<String>> listGroups = new LinkedList<>();
        final List<String> group = new LinkedList<>();
        stringStream.forEach(s -> {
            if (s.isBlank()) {
                listGroups.add(new LinkedList<>(group));
                group.clear();
            } else {
                group.add(s);
            }
        });
        if (!group.isEmpty()) {
            listGroups.add(group);
        }
        return listGroups;
    }
}
