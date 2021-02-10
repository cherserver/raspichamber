# raspichamber

Прошивка управления шкафом для 3D-принтера.
Шкаф сочетает в себе функции сушилки для филамента, обогреваемого короба для печати, вытяжки.

<table>
    <caption>Подключаемое к Raspberry Pi оборужование шкафа</caption>
    <tr>
        <th>Устройство</th>
        <th>Пин устройства</th>
        <th>Пин Raspberry GPIO</th>
    </tr>
    <tr>
        <td>Темп. датчик в сушилке</td>
        <td>DHT22 out</td>
        <td>GPIO2 (SDA)</td>
    </tr>
    <tr>
        <td>Темп. датчик в шкафу</td>
        <td>DHT22 out</td>
        <td>GPIO22</td>
    </tr>
    <tr>
        <td>Темп. датчик снаружи</td>
        <td>DHT22 out</td>
        <td>GPIO23</td>
    </tr>
    <tr>
        <td rowspan="4">Нагреватель</td>
        <td>Button 1</td>
        <td>GPIO26</td>
    </tr>
    <tr>
        <td>Button 2</td>
        <td>GPIO16</td>
    </tr>
    <tr>
        <td>Button 3</td>
        <td>GPIO20 (PCM_DIN)</td>
    </tr>
    <tr>
        <td>Button 4</td>
        <td>GPIO21 (PCM_DOUT)</td>
    </tr>
    <tr>
        <td rowspan="2">Вытяжной вентилятор</td>
        <td>Fan PWM in</td>
        <td>GPIO13 (PWM1)</td>
    </tr>
    <tr>
        <td>Fan Tachometer out</td>
        <td>GPIO4 (GPCLK0)</td>
    </tr>
    <tr>
        <td rowspan="2">Перепускной вентилятор</td>
        <td>Fan PWM in</td>
        <td>GPIO12 (PWM0)</td>
    </tr>
    <tr>
        <td>Fan Tachometer out</td>
        <td>GPIO17</td>
    </tr>
    <tr>
        <td rowspan="2">Вентилятор охлаждения RPi</td>
        <td>Fan PWM in</td>
        <td>GPIO0 (ID_SD)</td>
    </tr>
    <tr>
        <td>Fan Tachometer out</td>
        <td>GPIO1 (ID_SC)</td>
    </tr>
    <tr>
        <td>Клапан сушилки</td>
        <td>Servo in</td>
        <td>GPIO3 (SCL)</td>
    </tr>
</table>