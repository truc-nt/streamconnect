package com.hcmut.streamconnect.util;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.Set;
import java.util.StringJoiner;
import java.util.TreeSet;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.function.Supplier;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import java.util.stream.StreamSupport;

public class CollectionUtils {
    private CollectionUtils() {}

    public static <T> Collection<List<T>> chunk(Collection<T> collection, int chunkSize) {
        AtomicInteger counter = new AtomicInteger();
        return collection.stream()
                .collect(Collectors.groupingBy(e -> counter.getAndIncrement() / chunkSize))
                .values();
    }

    public static <T> Stream<T> streamOfNullable(Collection<T> c) {
        return Optional.ofNullable(c).stream().flatMap(Collection::stream);
    }

    public static <T, R> Stream<R> flatMap(Collection<T> c, Function<T, Collection<R>> fn) {
        if (c == null) return Stream.empty();
        return c.stream().flatMap(item -> fn.apply(item).stream());
    }

    public static <T, R> List<R> mapToList(Collection<T> c, Function<T, R> fn) {
        return streamOfNullable(c).map(fn).collect(Collectors.toList());
    }

    public static <T, R> Set<R> mapToSet(Collection<T> c, Function<T, R> fn) {
        return streamOfNullable(c).map(fn).collect(Collectors.toSet());
    }

    public static <T, R, C extends Collection<R>> C mapToCollection(Collection<T> c, Function<T, R> fn, Supplier<C> collectionFactory) {
        return streamOfNullable(c).map(fn).collect(Collectors.toCollection(collectionFactory));
    }

    public static <T> Set<String> mapToCaseInsensitiveSet(Collection<T> c, Function<T, String> fn) {
        return streamOfNullable(c).map(fn).collect(Collectors.toCollection(() -> new TreeSet<>(String.CASE_INSENSITIVE_ORDER)));
    }

    public static <T> List<T> filter(Collection<T> c, Predicate<T> predicate) {
        if (c == null) return Collections.emptyList();
        return c.stream().filter(predicate).collect(Collectors.toList());
    }

    public static <T> Optional<T> findFirst(Collection<T> c, Predicate<T> predicate) {
        if (c == null) return Optional.empty();
        return c.stream().filter(predicate).findFirst();
    }

    @SafeVarargs
    public static <T> boolean containsAny(Collection<? extends T> c, T... values) {
        if (c == null) return false;
        for (T value : values) {
            if (c.contains(value)) return true;
        }
        return false;
    }

    @SafeVarargs
    public static <T> boolean containsAnyOther(Collection<? extends T> c, T... values) {
        Set<T> others = removeAll(c, values);
        return !others.isEmpty();
    }

    @SafeVarargs
    public static <T> Set<T> removeAll(Collection<? extends T> c, T... values) {
        if (c == null) return Collections.emptySet();
        HashSet<T> set = new HashSet<>(c);
        set.removeAll(Arrays.asList(values));
        return set;
    }

    public static <T> List<T> newListWithDefault(int size, T defaultValue) {
        List<T> list = new ArrayList<>(size);
        for (int i = 0; i < size; i++) {
            list.add(defaultValue);
        }
        return list;
    }

    public static <T, K> Map<K, List<T>> groupingBy(Collection<T> c, Function<T, K> keyFunc) {
        if (c == null) return Collections.emptyMap();
        return c.stream().collect(Collectors.groupingBy(keyFunc));
    }

    public static <T, K> Map<K, Integer> groupingCount(Collection<T> c, Function<T, K> keyFunc) {
        if (c == null) return Collections.emptyMap();
        return c.stream().collect(Collectors.toMap(keyFunc, obj -> 1, Integer::sum));
    }

    public static <T extends V, K, V> Map<K, V> toMapByKey(Collection<T> c, Function<T, K> keyFunc) {
        if (c == null) return Collections.emptyMap();
        return c.stream().collect(Collectors.toMap(keyFunc, obj -> obj));
    }

    public static <T extends V, K, V> Map<K, V> toMapByKeySkipDuplicate(Collection<T> c, Function<T, K> keyFunc) {
        if (c == null) return Collections.emptyMap();
        return c.stream().collect(Collectors.toMap(keyFunc, obj -> obj, (o1, o2) -> o1));
    }

    public static <T, U, R> Map<T, R> mapEntryValue(Map<T, U> m, Function<U, R> fn) {
        if (m == null) return Collections.emptyMap();
        Map<T, R> result = new HashMap<>();
        m.forEach((k, v) -> result.put(k, fn.apply(v)));
        return result;
    }

    public static String join(Iterable<String> elements) {
        return join(elements, null);
    }

    public static String join(String delimiter, Iterable<String> elements) {
        return join(delimiter, elements, "", "");
    }

    public static String join(String delimiter, Iterable<String> elements, String prefix, String suffix) {
        return join(delimiter, elements, prefix, suffix, Function.identity());
    }

    public static <T> String join(Iterable<T> elements, Function<T, String> func) {
        return join(", ", elements, "", "", func);
    }

    public static <T> String join(String delimiter, Iterable<T> elements, String prefix, String suffix, Function<T, String> func) {
        final StringJoiner stringJoiner = new StringJoiner(delimiter, prefix, suffix);
        elements.forEach(ele -> stringJoiner.add(func != null ? func.apply(ele) : ele.toString()));
        return stringJoiner.toString();
    }

    public static <T> List<T> toList(Iterable<T> iterable) {
        return StreamSupport.stream(iterable.spliterator(), false).collect(Collectors.toList());
    }
}

