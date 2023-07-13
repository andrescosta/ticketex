import { Injectable, Logger } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { catchError, firstValueFrom } from 'rxjs';
import { AxiosError } from 'axios';

@Injectable()
export class RpcClient<T> {
  private readonly logger = new Logger(RpcClient.name);
  constructor(
    protected readonly httpService: HttpService,
    protected readonly url: string,
  ) {}

  async post(t: T): Promise<T> {
    return t;
  }
  async patch(t: T): Promise<T> {
    return t;
  }
  async get(): Promise<T[]> {
    const { data } = await firstValueFrom(
      this.httpService.get<T[]>(this.url).pipe(
        catchError((error: AxiosError) => {
          this.logger.error(error.response.data);
          throw 'An error happened!';
        }),
      ),
    );
    return data;
  }
  async getById(id: string): Promise<T> {
    const { data } = await firstValueFrom(
      this.httpService.get<T>(this.url + id).pipe(
        catchError((error: AxiosError) => {
          this.logger.error(error.response.data);
          throw 'An error happened!';
        }),
      ),
    );
    return data;
  }
}
