

*** Settings ***
Documentation    Smoketest Flink execution
Library          OperatingSystem
Library          String

*** Test Cases ***

Check version
  ${output} =      Execute                     flink --version

Run Sample WordCount
  ${output} =      Execute                     flink run /opt/flink/examples/batch/WordCount.jar
                   Should Contain              ${output}           something

*** Keywords
Execute
    [arguments]                     ${command}
    ${rc}                           ${output} =                 Run And Return Rc And Output           ${command}
    Log                             ${output}
    Should Be Equal As Integers     ${rc}                       0
    [return]                        ${output}
