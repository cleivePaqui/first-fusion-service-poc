#!groovy

// This 'import' happens automatically but is a useful reference if you want to use a feature branch instead of master/main
@Library(['msp@master', 'fusion@main']) _

def devConfig = [:]
def stgConfig = [:]
def prodConfig = [:]
def prodRegions = []

def kanikoRegistries = [
    "API_KANIKO_REGISTRY": [
        "cmd/api-server/Dockerfile",
        "616350704275.dkr.ecr.eu-west-1.amazonaws.com/fusion/api-template-server-snapshot",
        "616350704275.dkr.ecr.eu-west-1.amazonaws.com/fusion/api-template-server"
    ]
]

fusionCloudPipeline(
    "fusion-api-template",
    "Workstream 2",
    "FSN-11581",
    devConfig,
    stgConfig,
    prodConfig,
    prodRegions,
    kanikoRegistries
)
