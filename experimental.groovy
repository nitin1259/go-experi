def artifactRegexList = ["doedemo*.war", "doedemo*.jar"]
def artifactPaths = [
    "com/example/cloud/dctrn-xyz/2024.06.02/dctrnDemo-2021.04.31.45.jar",
    "com/example/cloud/demodoeapp/2024.06.02/dctrn-2024.06.01.45.war",
    "com/example/cloud/dctrn-abc/2024.06.02/dctrn-2024.06.01.45.war",
    "com/example/cloud/demodoeapp/2024.06.01/dctrn-2024.06.01.45.module"
]

// Compile regex patterns once
def compiledPatterns = artifactRegexList.collect { pattern ->
    ~pattern.replace('*', '.*')
}

artifactPaths.each { path ->
    boolean matches = compiledPatterns.any { regex ->
        path ==~ regex
    }
    if (matches) {
        println "Processing artifact: ${path}"
    } else {
        println "Skipping artifact: ${path}"
    }
}

//--------

def artifactRegexList = ["doedemo*.war", "doedemo*.jar"]
def artifactPaths = [
    "com/example/cloud/dctrn-xyz/2024.06.02/dctrnDemo-2021.04.31.45.jar",
    "com/example/cloud/demodoeapp/2024.06.02/dctrn-2024.06.01.45.war",
    "com/example/cloud/dctrn-abc/2024.06.02/deodemodctrn-2024.06.01.45.war",
    "com/example/cloud/demodoeapp/2024.06.01/dctrn-2024.06.01.45.module"
]

// Compile regex patterns once
def compiledPatterns = artifactRegexList.collect { pattern ->
    ~pattern.replace('*', '.*')
}

artifactPaths.each { path ->
    boolean matches = compiledPatterns.any { pattern ->
        path ==~ /.*\/${pattern}/
    }
    if (matches) {
        println "Processing artifact: ${path}"
    } else {
        println "Skipping artifact: ${path}"
    }
}
