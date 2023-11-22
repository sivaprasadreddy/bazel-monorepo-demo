load("@rules_jvm_external//:defs.bzl", "artifact")
load("@contrib_rules_jvm//java:defs.bzl", "JUNIT5_DEPS", "java_test_suite")

java_library(
    name = "greeter",
    srcs = glob(["src/main/java/**/*.java"]),
)

java_test_suite(
    name = "greeter-tests",
    srcs = glob(["src/test/java/**/*.java"]),
    runner = "junit5",
    test_suffixes = [
        "Test.java",
        "Tests.java",
    ],
    runtime_deps = JUNIT5_DEPS,
    deps = [
        ":greeter",
        artifact("org.junit.jupiter:junit-jupiter-api"),
        artifact("org.junit.jupiter:junit-jupiter-params"),
    ],
)
