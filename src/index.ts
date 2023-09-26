import { Elysia } from "elysia";
import { jwt } from "@elysiajs/jwt";

const app = new Elysia().use(
  jwt({
    name: "jwt",
    secret: "VINO_JWT",
  }),
);

app.post("/password/login", async ({ body, jwt }) => {
  if (!body) {
    return {
      code: 1,
      message: "不合法的用户名或密码",
      data: null,
    };
  }
  const jwtStr = await jwt.sign(body as Record<string, string>);
  return {
    code: 0,
    message: "登录成功",
    data: {
      accessToken: jwtStr,
      refreshToken: jwtStr,
      expiresIn: 10,
    },
  };
});

app.get("/user", async ({ jwt, headers }) => {
  try {
    const jwtStr = headers.authorization?.split(" ")[1];
    console.log(jwtStr)
    const userInfo = (await jwt.verify(jwtStr)) as { username: string };
    if (userInfo) {
      return {
        userId: "fake_user_id",
        username: userInfo.username,
      };
    } else {
      return {
        code: 1,
        message: "错误的token",
        data: null,
      };
    }
  } catch (e) {
    return {
      code: 1,
      message: "JWT 解析错误",
      data: null,
    };
  }
});

app.listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);
