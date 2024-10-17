import {
  CallHandler,
  ConsoleLogger,
  ExecutionContext,
  Injectable,
  NestInterceptor,
} from '@nestjs/common';
import { yellow } from 'colors';
import { Request } from 'express';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
// import { RequestWithUser } from 'src/resources/guards/auth.guard';

@Injectable()
export class LoggerInterceptor implements NestInterceptor {
  constructor(private readonly logger: ConsoleLogger) {}

  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    const httpContext = context.switchToHttp();
    const request = httpContext.getRequest<Request>();
    const userAgent = request.get('user-agent') || '';
    const { ip } = request;
    const now = Date.now();
    return next.handle().pipe(
      tap(() => {
        this.logger.log(
          `${userAgent} ${ip}: ${request.method} ${request.url} ` +
            yellow(`+${Date.now() - now}ms`),
          'Logger',
        );
      }),
    );
  }
}
