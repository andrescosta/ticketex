import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { Adventure } from '../models/adventure';
import { AxiosResponse } from 'axios';
import { Observable } from 'rxjs';

@Injectable()
export class AdventuresService {
  constructor(private readonly httpService: HttpService) {}

  findAll(): Observable<AxiosResponse<Adventure[]>> {
    return this.httpService.get('http://localhost:3000/cats');
  }
}
