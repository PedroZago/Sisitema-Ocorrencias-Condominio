import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { ResponsavelModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class ResponsavelService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarResponsavel(responsavel: ResponsavelModel): Observable<any> {
    let params = `/api/responsavel/cadastrar`;
    return this.http.post(this.BASE_URL + params, responsavel, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodasUnidadesPorResponsavel(): Observable<any> {
    let params = `/api/responsavel/unidades`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  atualizarResponsavel(idUni: number, idUsu: number, responsavel: ResponsavelModel): Observable<any> {
    let params = `/api/responsavel/${idUni}/${idUsu}`;
    return this.http.put(this.BASE_URL + params, responsavel, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  deletarResponsavel(idUni: number, idUsu: number): Observable<any> {
    let params = `/api/responsavel/${idUni}/${idUsu}`;
    return this.http.delete(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }
}
