import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { ApartamentoModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class ApartamentosService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarApartamento(apartamento: ApartamentoModel): Observable<any> {
    let params = `/api/apartamento/cadastrar`;
    return this.http.post(this.BASE_URL + params, apartamento, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarApartamentoID(id: number): Observable<any> {
    let params = `/api/apartamento/${id}`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodosApartamentos(): Observable<any> {
    let params = `/api/apartamento`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  deletarApartamento(id: number): Observable<any> {
    let params = `/api/apartamento/${id}`;
    return this.http.delete(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  atualizarApartamento(id: number, apartamento: ApartamentoModel): Observable<any> {
    let params = `/api/apartamento/${id}`;
    return this.http.put(this.BASE_URL + params, apartamento, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }
}
