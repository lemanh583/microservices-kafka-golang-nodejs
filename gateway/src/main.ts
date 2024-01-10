import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { createProxyMiddleware } from 'http-proxy-middleware';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.use("/sign-up", createProxyMiddleware({
    target: "http://127.0.0.1:3001/",
    changeOrigin: true,
    
  }))
  await app.listen(3000);
}
bootstrap();
