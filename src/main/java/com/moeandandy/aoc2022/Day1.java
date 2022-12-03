package com.moeandandy.aoc2022;

import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Comparator;
import java.util.stream.Collectors;

public class Day1 {

    public static final String DEFAULT_FILE = "2022/day1.txt";

    public static void main(String[] args) throws URISyntaxException, IOException {
        String inputFilename = args.length > 0
            ? args[0]
            : DEFAULT_FILE;
        Path input = Paths.get(Day1.class.getClassLoader().getResource(inputFilename).toURI());
        System.out.println(StreamUtils.splitByBlanks(Files.lines(input))
            .stream()
            .map(group -> group
                .stream()
                .collect(Collectors.summingLong(Long::parseLong)))
            .sorted(Comparator.comparing(Long::longValue).reversed())
            .limit(3)
            .peek(System.out::println)
            .collect(Collectors.summingLong(Long::longValue)));
    }
}
