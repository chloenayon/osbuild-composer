#!/bin/bash
set -euo pipefail

# Get OS and architecture details.
source /etc/os-release
ARCH=$(uname -m)

WORKING_DIRECTORY=/usr/libexec/osbuild-composer
IMAGE_TEST_CASE_RUNNER=/usr/libexec/tests/osbuild-composer/osbuild-image-tests
IMAGE_TEST_CASES_PATH=/usr/share/tests/osbuild-composer/cases

PASSED_TESTS=()
FAILED_TESTS=()

# Print out a nice test divider so we know when tests stop and start.
test_divider () {
    printf "%0.s-" {1..78} && echo
}

# Get a list of test cases.
get_test_cases () {
    TEST_CASE_SELECTOR="${ID}_${VERSION_ID%.*}-${ARCH}*.json"
    pushd $IMAGE_TEST_CASES_PATH > /dev/null
        ls $TEST_CASE_SELECTOR
    popd > /dev/null
}

# Run a test case and store the result as passed or failed.
run_test_case () {
    TEST_RUNNER=$1
    TEST_CASE_FILENAME=$2
    TEST_NAME=$(basename $TEST_CASE_FILENAME)

    echo
    test_divider
    echo "🏃🏻 Running test: ${TEST_NAME}"
    test_divider

    # Set up the testing command.
    TEST_CMD="$TEST_RUNNER -test.v ${IMAGE_TEST_CASES_PATH}/${TEST_CASE_FILENAME}"

    # Run the test and add the test name to the list of passed or failed
    # tests depending on the result.
    if sudo $TEST_CMD 2>&1 | tee ${WORKSPACE}/${TEST_NAME}.log; then
        PASSED_TESTS+=("$TEST_NAME")
    else
        FAILED_TESTS+=("$TEST_NAME")
    fi

    test_divider
    echo
}

# Ensure osbuild-composer-tests is installed.
if ! rpm -qi osbuild-composer-tests > /dev/null 2>&1; then
    sudo dnf -y install osbuild-composer-tests
fi

# Change to the working directory.
cd $WORKING_DIRECTORY

# Run each test case.
for TEST_CASE in $(get_test_cases); do
    run_test_case $IMAGE_TEST_CASE_RUNNER $TEST_CASE
done

# Print a report of the test results.
test_divider
echo "😃 Passed tests: " "${PASSED_TESTS[@]}"
echo "☹ Failed tests: " "${FAILED_TESTS[@]}"
test_divider

# Exit with a failure if any tests failed.
if [ ${#FAILED_TESTS[@]} -eq 0 ]; then
    echo "🎉 All tests passed."
    exit 0
else
    echo "🔥 One or more tests failed."
    exit 1
fi