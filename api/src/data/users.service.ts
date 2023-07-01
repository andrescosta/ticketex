import { Injectable, Logger } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { User } from '../models/user';
import { catchError, firstValueFrom } from 'rxjs';
import { AxiosError } from 'axios';

@Injectable()
export class UsersService {
  private readonly logger = new Logger(UsersService.name);
  constructor(private readonly httpService: HttpService) {}

  async findAll(): Promise<User[]> {
    const { data } = await firstValueFrom(
      this.httpService.get<User[]>('http://localhost:8001/cats').pipe(
        catchError((error: AxiosError) => {
          this.logger.error(error.response.data);
          throw 'An error happened!';
        }),
      ),
    );
    return data;
  }
}
