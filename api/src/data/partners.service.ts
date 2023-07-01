import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { Partner } from '../models/partner';
import { AxiosResponse } from 'axios';
import { Observable } from 'rxjs';

@Injectable()
export class PartnersService {
  constructor(private readonly httpService: HttpService) {}

  findAll(): Observable<AxiosResponse<Partner[]>> {
    return this.httpService.get('http://localhost:3000/cats');
  }
}
