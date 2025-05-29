import { BadRequestException, ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationError } from 'class-validator';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {
    cors: {
      origin: 'http://localhost:3000',
      // allowedHeaders: '*',
      credentials: true,
      optionsSuccessStatus: 200,
      methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE'],
    },
  });

  app.useGlobalPipes(
    new ValidationPipe({
      exceptionFactory: (errors: ValidationError[]) => {
        const errorsMap: Record<string, string> = {};
        for (const error of errors) {
          errorsMap[error.property] = Object.values(
            error?.constraints ?? {},
          )[0];
        }
        return new BadRequestException({
          success: false,
          message: 'validation error',
          errors: errorsMap,
        });
      },
    }),
  );
  await app.listen(process.env.PORT ?? 8000);
}
bootstrap();
