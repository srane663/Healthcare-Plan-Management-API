package service

import (
    "assignment1/redis"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "github.com/gin-gonic/gin"
)

func GenerateETag(plan interface{}) string {

    bytes, _ := json.Marshal(plan)

    hash := sha256.Sum256(bytes)

    return base64.StdEncoding.EncodeToString(hash[:])
}

func CreatePlan(plan map[string]interface{}) (int, string, interface{}) {

    objectID := plan["objectId"].(string)

    exists, _ := redis.Client.Exists(redis.Ctx, objectID).Result()

    if exists == 1 {
        return 409, "", gin.H{
            "error": "Plan already exists",
        }
    }

    etag := GenerateETag(plan)

    plan["etag"] = etag

    bytes, _ := json.Marshal(plan)

    redis.Client.Set(
        redis.Ctx,
        objectID,
        bytes,
        0,
    )

    return 201, etag, plan
}

func GetPlan(
    id string,
    clientETag string,
) (int, string, interface{}) {

    result, err := redis.Client.Get(
        redis.Ctx,
        id,
    ).Result()

    if err != nil {
        return 404, "", gin.H{
            "error": "Plan not found",
        }
    }

    var plan map[string]interface{}

    json.Unmarshal(
        []byte(result),
        &plan,
    )

    etag := plan["etag"].(string)

    if clientETag != "" &&
        clientETag == etag {

        return 304, etag, nil
    }

    return 200, etag, plan
}

func DeletePlan(id string) int {

    exists, _ := redis.Client.Exists(
        redis.Ctx,
        id,
    ).Result()

    if exists == 0 {
        return 404
    }

    redis.Client.Del(
        redis.Ctx,
        id,
    )

    return 204
}
