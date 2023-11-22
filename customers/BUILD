load("@rules_jvm_external//:defs.bzl", "artifact")
load("@contrib_rules_jvm//java:defs.bzl", "JUNIT5_DEPS", "java_test_suite")

java_library(
    name = "customers-lib",
    srcs = glob(["src/main/java/**/*.java"]),
    deps = [
        artifact("org.postgresql:postgresql"),
        artifact("ch.qos.logback:logback-classic"),
    ],
)

java_test_suite(
    name = "customers-lib-tests",
    srcs = glob(["src/test/java/**/*.java"]),
    runner = "junit5",
    test_suffixes = [
        "Test.java",
        "Tests.java",
    ],
    runtime_deps = JUNIT5_DEPS,
    deps = [
        ":customers-lib",
        artifact("org.junit.jupiter:junit-jupiter-api"),
        artifact("org.junit.jupiter:junit-jupiter-params"),
        artifact("org.testcontainers:postgresql"),
    ],
)
