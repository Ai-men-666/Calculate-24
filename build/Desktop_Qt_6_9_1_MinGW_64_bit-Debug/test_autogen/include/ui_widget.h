/********************************************************************************
** Form generated from reading UI file 'widget.ui'
**
** Created by: Qt User Interface Compiler version 6.9.1
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_WIDGET_H
#define UI_WIDGET_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_Widget
{
public:
    QPushButton *Confirm;
    QPushButton *Exit;
    QLineEdit *Card1;
    QLineEdit *Card2;
    QLineEdit *Card4;
    QLineEdit *Card3;
    QLabel *label;
    QLabel *label_2;
    QLabel *label_3;
    QLabel *res;

    void setupUi(QWidget *Widget)
    {
        if (Widget->objectName().isEmpty())
            Widget->setObjectName("Widget");
        Widget->resize(424, 272);
        Confirm = new QPushButton(Widget);
        Confirm->setObjectName("Confirm");
        Confirm->setGeometry(QRect(70, 210, 81, 41));
        QFont font;
        font.setPointSize(18);
        Confirm->setFont(font);
        Exit = new QPushButton(Widget);
        Exit->setObjectName("Exit");
        Exit->setGeometry(QRect(250, 210, 91, 41));
        Exit->setFont(font);
        Card1 = new QLineEdit(Widget);
        Card1->setObjectName("Card1");
        Card1->setGeometry(QRect(140, 20, 113, 20));
        QFont font1;
        font1.setPointSize(9);
        Card1->setFont(font1);
        Card2 = new QLineEdit(Widget);
        Card2->setObjectName("Card2");
        Card2->setGeometry(QRect(140, 50, 113, 20));
        Card4 = new QLineEdit(Widget);
        Card4->setObjectName("Card4");
        Card4->setGeometry(QRect(140, 110, 113, 20));
        Card3 = new QLineEdit(Widget);
        Card3->setObjectName("Card3");
        Card3->setGeometry(QRect(140, 80, 113, 20));
        label = new QLabel(Widget);
        label->setObjectName("label");
        label->setGeometry(QRect(50, 20, 91, 31));
        QFont font2;
        font2.setPointSize(17);
        label->setFont(font2);
        label_2 = new QLabel(Widget);
        label_2->setObjectName("label_2");
        label_2->setGeometry(QRect(30, 60, 71, 41));
        QFont font3;
        font3.setPointSize(19);
        label_2->setFont(font3);
        label_3 = new QLabel(Widget);
        label_3->setObjectName("label_3");
        label_3->setGeometry(QRect(70, 150, 51, 31));
        QFont font4;
        font4.setPointSize(16);
        label_3->setFont(font4);
        res = new QLabel(Widget);
        res->setObjectName("res");
        res->setGeometry(QRect(140, 160, 91, 41));

        retranslateUi(Widget);

        QMetaObject::connectSlotsByName(Widget);
    } // setupUi

    void retranslateUi(QWidget *Widget)
    {
        Widget->setWindowTitle(QCoreApplication::translate("Widget", "Test", nullptr));
        Confirm->setText(QCoreApplication::translate("Widget", "\347\241\256\345\256\232", nullptr));
        Exit->setText(QCoreApplication::translate("Widget", "\351\200\200\345\207\272", nullptr));
        Card1->setText(QString());
        label->setText(QCoreApplication::translate("Widget", "\350\276\223\345\205\245\345\215\241\347\211\214\345\200\274", nullptr));
        label_2->setText(QCoreApplication::translate("Widget", "\347\256\22724\347\202\271", nullptr));
        label_3->setText(QCoreApplication::translate("Widget", "\347\273\223\346\236\234", nullptr));
        res->setText(QString());
    } // retranslateUi

};

namespace Ui {
    class Widget: public Ui_Widget {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_WIDGET_H
