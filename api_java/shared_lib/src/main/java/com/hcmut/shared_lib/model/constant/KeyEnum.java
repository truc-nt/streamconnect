package com.hcmut.shared_lib.model.constant;

import java.util.EnumSet;
import java.util.HashMap;
import java.util.Map;

public interface KeyEnum<E extends Enum<E> & KeyEnum> {
    String getKey();
    default E get(Class<E> enumType, String key) {
        return KeyEnumUtils.get(enumType, key);
    }

    class KeyEnumUtils {
        private static Map<Class, Map> reverseLookupMaps = new HashMap<>();

        public static <E extends Enum<E> & KeyEnum> E get(Class<E> enumType, String key) {
            if (!reverseLookupMaps.containsKey(enumType)) {
                Map<String, E> reverse_lookup = new HashMap<>();
                for (E metric : EnumSet.allOf(enumType)) {
                    reverse_lookup.put(metric.getKey(), metric);
                }
                reverseLookupMaps.put(enumType, reverse_lookup);
            }
            //noinspection unchecked
            return (E) reverseLookupMaps.get(enumType).get(key);
        }
    }
}
