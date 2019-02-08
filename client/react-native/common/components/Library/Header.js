import { View } from 'react-native'
import React, { PureComponent } from 'react'
import { Button, Flex, Text, SearchBar } from '.'
import { colors } from '../../constants'
import { padding, borderBottom, paddingBottom } from '../../styles'
import { isRTL } from '../../i18n'

const [defaultTextColor, defaultBackColor] = [colors.black, colors.white]

const HeaderButton = ({ icon, color, style, ...otherProps }) => {
  return <Button icon={icon} large color={color} {...otherProps} />
}

export default class Header extends PureComponent {
  render () {
    const {
      navigation,
      title,
      titleIcon,
      backBtn,
      rightBtn,
      rightBtnIcon,
      onPressRightBtn,
      searchBar,
      searchHandler,
    } = this.props

    const colorText =
      this.props.colorText == null ? defaultTextColor : this.props.colorText
    const colorBack =
      this.props.colorBack == null ? defaultBackColor : this.props.colorBack
    const colorBtnLeft =
      this.props.colorBtnLeft == null
        ? defaultTextColor
        : this.props.colorBtnLeft
    const colorBtnRight =
      this.props.colorBtnRight == null
        ? defaultTextColor
        : this.props.colorBtnRight

    let searchBarComponent = null
    if (searchBar === true) {
      searchBarComponent = (
        <SearchBar onChangeText={text => searchHandler(text)} />
      )
    } else if (searchBar !== undefined && searchBar !== false) {
      searchBarComponent = searchBar
    }

    // todo calc height in function of dev status bar and search bar
    // let height

    return (
      <View
        style={[
          {
            backgroundColor: colorBack,
            height: searchBar && true ? 110 : 70,
          },
          borderBottom,
          padding,
        ]}
      >
        <Flex.Rows>
          <Flex.Cols
            size={1}
            justify='between'
            align='center'
            style={[searchBar ? paddingBottom : {}]}
          >
            {backBtn && (
              <HeaderButton
                icon='arrow-left'
                color={colorBtnLeft}
                onPress={() => {
                  if (typeof backBtn === 'function') {
                    backBtn()
                  }
                  navigation.goBack(null)
                }}
                flip={isRTL()}
                justify='start'
                middle
              />
            )}
            <Text
              icon={titleIcon}
              left
              large
              color={colorText}
              justify={backBtn ? 'center' : 'start'}
              middle
              size={5}
            >
              {title}
            </Text>
            {rightBtn ? <View>{rightBtn}</View> : null}
            {!rightBtn &&
            rightBtnIcon !== null && (
              <HeaderButton
                icon={rightBtnIcon}
                color={colorBtnRight}
                onPress={onPressRightBtn}
                justify='end'
                middle
              />
            )}
          </Flex.Cols>
          {searchBarComponent}
        </Flex.Rows>
      </View>
    )
  }
}

Header.HeaderButton = HeaderButton
