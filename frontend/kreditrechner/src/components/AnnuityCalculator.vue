<template>
    <v-app id="inspire">
      <v-app-bar
        class="px-3"
        color="white"
        flat
        density="compact"
      >
        <v-avatar
          color="grey-darken-1"
          size="32"
        ></v-avatar>
  
        <v-spacer></v-spacer>
  
        <v-tabs
          centered
          color="grey-darken-2"
        >
          <v-tab
            v-for="link in links"
            :key="link"
          >
            {{ link }}
          </v-tab>
        </v-tabs>
        <v-spacer></v-spacer>
      </v-app-bar>
  
      <v-main class="bg-grey-lighten-3">
        <v-container>
          <v-row>  
            <v-col
              cols="12"
              sm="8"
            >
              <v-sheet
                class="pa-12"
                min-height="70vh"
                rounded="lg"
              >
              <!-- middle col content  -->
                <v-sheet>
                  <v-row>
                    <v-icon color="blue darken-2" icon="mdi-calculator"></v-icon>
                    <h2 class="text-blue darken-2">Tilgungsrechner</h2>
                  </v-row>
                </v-sheet>
                <v-sheet
                class="py-12"
                >
                    <v-text-field
                    v-model="creditsum"
                    clearable
                    label="Darlehensbetrag"
                    type="text"
                    variant="outlined"
                    :rules="creditsumRules"
                    :counter="10"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihres Darlehens ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-currency-eur"></v-icon>
                      </template>
                    </v-text-field>

                    <v-text-field
                    v-model="interestRate"
                    clearable
                    label="Sollzinssatz"
                    type="text"
                    variant="outlined"
                    :rules="rateRules"
                    :counter="4"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihres Sinssatzes ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-percent"></v-icon>
                      </template>
                    </v-text-field>
                    <v-text-field
                    v-model="repaymentRate"
                    clearable
                    label="Tilgungssatz"
                    type="text"
                    variant="outlined"
                    :rules="rateRules"
                    :counter="4"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihrer anfänglichen geplanten Tilgung ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-percent"></v-icon>
                      </template>
                    </v-text-field>
                    <v-text-field
                    v-model="unscheduledRepaymentRate"
                    clearable
                    label="Sondertilgung"
                    type="text"
                    variant="outlined"
                    :rules="rateRules"
                    :counter="4"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihrer jährlichen geplanten Sondertilgung ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-percent"></v-icon>
                      </template>
                    </v-text-field>
                    <v-text-field
                    v-model="runtime"
                    clearable
                    label="Zinsbindungsdauer"
                    type="text"
                    variant="outlined"
                    :rules="runtimeRules"
                    :counter="2"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Dauer in Jahren der Zinsbindung ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-av-timer"></v-icon>
                      </template>
                    </v-text-field>
                    <v-btn
                  color="success"
                  class="mr-4"
                  @click="calculateAnnuity"
                >
                  Berechnen
                </v-btn>

                <v-sheet
                class="py-4" 
                v-if="show_calculation">
                  <v-divider></v-divider>
                  <v-row class="py-6">
                    <v-col
                    cols="12"
                    sm="6">
                    <div class="text-body">
                      Monatliche Rate
                    </div>
                  </v-col>
                  <v-col
                  cols="12"
                  sm="6">
                    <div class="text-body text-right">
                      {{number_formatter.format(
                        roundOff((calculatedAnnuity.results[0].annuity / 12), 2))
                      }}
                    </div>
                  </v-col>
                  </v-row>
                  
                  <v-table
                  density="compact">
                  <thead>
                    <tr>
                      <th class="text-left">
                        Jahr
                      </th>
                      <th class="text-left">
                        Restschuld
                      </th>
                      <th class="text-left">
                        Zinsanteil
                      </th>
                      <th class="text-left">
                        Tilgung
                      </th>
                      <th class="text-left">
                        Sondertilgung
                      </th>
                      <th class="text-left">
                        Rate
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="item in calculatedAnnuity.results">
                      <td>{{item.year}}</td>
                      <td>{{number_formatter.format(item.residualDept)}}</td>
                      <td>{{number_formatter.format(item.interest)}}</td>
                      <td>{{number_formatter.format(item.repayment)}}</td>
                      <td>{{number_formatter.format(item.unscheduledRepayment)}}</td>
                      <td>{{number_formatter.format(item.annuity)}}</td>
                    </tr>
                  </tbody>
                </v-table>

                <v-card
                    max-width="400"
                    class="mx-auto"
                  >
                    <v-container>
                    <v-row dense>
                      <v-col cols="12">
                        <v-card
                          color="primary"
                          theme="light"
                        >
                          <v-card-title class="text-h6">
                            Zusammenfassung
                          </v-card-title>
                          <v-card-subtitle>Über die Laufzeit insgesamt gezahlte Positionen</v-card-subtitle>
                          <v-divider thickness="2"></v-divider>
                          <v-card-text>
                            <v-row>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text--primary">Zinsen</p>
                              </v-col>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text-right">{{number_formatter.format(calculatedAnnuity.interestSum)}}</p>
                              </v-col>
                          </v-row>
                          <v-row>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text--primary">Raten</p>
                              </v-col>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text-right">{{number_formatter.format(calculatedAnnuity.annuitySum)}}</p>
                              </v-col>
                          </v-row>
                          <v-row>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text--primary">Tilgung</p>
                              </v-col>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text-right">{{number_formatter.format(calculatedAnnuity.repaymentSum)}}</p>
                              </v-col>
                          </v-row>
                          <v-row>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text--primary">Sondertilgung</p>
                              </v-col>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text-right">{{number_formatter.format(calculatedAnnuity.unscheduledRepaymentSum)}}</p>
                              </v-col>
                          </v-row>
                          <v-row>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text--primary">Insgesamt eingezahlt</p>
                              </v-col>
                              <v-col
                              cols="12"
                              sm="6">
                              <p class="text-right">{{number_formatter.format(calculatedAnnuity.totalPayment)}}</p>
                              </v-col>
                          </v-row>
                          </v-card-text>
                          <v-divider></v-divider>
                          <v-card-actions>
                            <v-btn variant="text">
                              Export CSV
                            </v-btn>
                          </v-card-actions>
                        </v-card>
                      </v-col>
                    </v-row>
                    </v-container>
                  </v-card>
                </v-sheet>
                </v-sheet>
                </v-sheet>
            </v-col>
  
            <v-col
              cols="12"
              sm="2"
            >
              <v-sheet
                rounded="lg"
                min-height="268"
              >
                <!--  -->
              </v-sheet>
            </v-col>
          </v-row>
        </v-container>
      </v-main>
    </v-app>
  </template>
  
  <script lang="ts">
  import axios from 'axios';

  class AnnuityResult {
    public year: number;
    public residualDept: number;
    public interest: number;
    public repayment: number;
    public unscheduledRepayment: number;
    public annuity: number;

    public constructor() {
      this.year = 0;
      this.residualDept = 0;
      this.interest = 0;
      this.repayment = 0;
      this.unscheduledRepayment = 0;
      this.annuity = 0;
    }
  }

  class AnnuityResponse {
    public results: AnnuityResult[];
    public interestSum: number;
    public repaymentSum: number;
    public unscheduledRepaymentSum: number;
    public annuitySum: number;
    public totalPayment: number;

    public constructor() {
      this.results = [];
      this.interestSum = 0;
      this.repaymentSum = 0;
      this.unscheduledRepaymentSum = 0;
      this.annuitySum = 0;
      this.totalPayment = 0;
    }
  }

  class AnnuityRequest {
      public creditsum: number;
      public runtime: number;
      public interestRate: number;
      public initialRepaymentRate: number;
      public unscheduledRepaymentRate: number;

      constructor() {
        this.creditsum = 0;
        this.runtime = 0;
        this.interestRate = 0;
        this.initialRepaymentRate = 0;
        this.unscheduledRepaymentRate = 0;
      }
  }

    export default {
      data: () => ({
        calculatedAnnuity: new AnnuityResponse(),
        show_calculation: false,
        interestRate:'2,5',
        repaymentRate:'2,0',
        unscheduledRepaymentRate:'1,0',
        runtime:'10',
        creditsum: '100000',
        number: 100000,
        number_formatter: new Intl.NumberFormat('de-DE', { style: 'currency', currency: 'EUR' }),

        links: [
          'Tilgungsrechner',
          'Bausparrechner'
        ],
        creditsumRules: [
          (v: string | any[]) => !!v || 'Creditsum is required',
          (v: string | any[]) => v.length <= 10 || 'Creditsum must be less than 10 characters',
        ],
        runtimeRules: [
          (v: any) => !!v || 'Runtime is required',
          (v: string | any[]) => v.length <= 2 || 'Runtime must be less than 3 characters',
        ],
        rateRules: [
          (v: any) => !!v || 'Rate is required',
          (v: string | any[]) => v.length <= 4 || 'Rate must be less than 5 characters',
          (v: string) => {
          const pattern_float = /^[0-9]\,[0-9]+$/;
          const pattern_int = /^[0-9]+$/;
          return  pattern_float.test(v) ||
                  pattern_int.test(v) ||
                  'Invalid rate.'
        },
        ],
      }),
      methods: {
        roundOff(num: number, places: number): number {
          const x = Math.pow(10,places);
          return Math.round(num * x) / x;
        },
        convertToAnnuityRequest(): AnnuityRequest {
          var annuityRequest = new AnnuityRequest();
          annuityRequest.creditsum = parseInt(this.creditsum);
          annuityRequest.interestRate = parseFloat(this.interestRate.replace(',','.'));
          annuityRequest.initialRepaymentRate = parseFloat(this.repaymentRate.replace(',','.'));
          annuityRequest.unscheduledRepaymentRate = parseFloat(this.unscheduledRepaymentRate.replace(',','.'));
          annuityRequest.runtime = parseInt(this.runtime);

          return annuityRequest;
        },
        async calculateAnnuity() {
          this.show_calculation = false;
          var annuityRequest = this.convertToAnnuityRequest();
          var url = window.location.href + "api/v1/calculateAnnuity"
          const {data, status} = await axios.post<AnnuityResponse>(
            url,
            JSON.stringify(annuityRequest),
            {
              headers: {
                ContentType: 'application/json'
              }
            }
          );
          
          if(status != 200) {
            console.log("http status error: ", status);
            return;
          }

          this.calculatedAnnuity = data
          if(this.calculatedAnnuity.results.length <= 0) {
            console.log("invalid annuity: ", this.calculatedAnnuity)
            return;
          }
          this.show_calculation = true;
        }
      }
    }
  </script>